package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"webapp/server/config"
	"webapp/server/database"
	"webapp/server/models"
	"webapp/server/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"
)

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
}

func googleOAuthConfig() *oauth2.Config {
	cfg := config.LoadConfig()
	return &oauth2.Config{
		ClientID:     cfg.GoogleClientID,
		ClientSecret: cfg.GoogleClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  "postmessage",
	}
}

func GoogleLogin(ctx *gin.Context) {
	oauthConfig := googleOAuthConfig()
	url := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	ctx.JSON(http.StatusOK, gin.H{"url": url})
}

func GoogleCallback(ctx *gin.Context) {
	code := ctx.Query("code")
	if code == "" {
		err := errors.New("code not provided")
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	oauthConfig := googleOAuthConfig()
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client := oauthConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer response.Body.Close()

	userData, err := io.ReadAll(response.Body)
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userInfo GoogleUserInfo
	if err := json.Unmarshal(userData, &userInfo); err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()

	var user models.User
	result := db.Where("google_id = ?", userInfo.ID).First(&user)

	user.GoogleID = userInfo.ID
	user.Name = userInfo.Name
	user.Email = userInfo.Email

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Create new user
			if err := db.Create(&user).Error; err != nil {
				glog.Error(err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			glog.Error(result.Error)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
	} else {
		// Update existing user
		if err := db.Save(&user).Error; err != nil {
			glog.Error(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	jwtToken, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":   jwtToken,
		"user_id": fmt.Sprintf("%d", user.ID),
		"name":    user.Name,
		"email":   user.Email,
	})
}

func shouldRegister(result *gorm.DB) (bool, int) {
	if result.Error == nil {
		if result.RowsAffected == 0 {
			return true, http.StatusNotFound
		} else {
			return false, http.StatusFound
		}
	} else {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return true, http.StatusNotFound
		} else {
			return false, http.StatusInternalServerError
		}
	}
}

func shouldLogin(result *gorm.DB) (bool, int) {
	if result.Error == nil {
		if result.RowsAffected != 0 {
			return true, http.StatusFound
		} else {
			return false, http.StatusUnauthorized
		}
	} else {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, http.StatusNotFound
		} else {
			return false, http.StatusInternalServerError
		}
	}
}

func Register(ctx *gin.Context) {
	var registerRequest models.RegisterRequest

	if err := ctx.BindJSON(&registerRequest); err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if registerRequest.Email == "" || registerRequest.Password == "" || registerRequest.Name == "" {
		err := errors.New("missing required fields")
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if registerRequest.Password != registerRequest.ConfirmPassword {
		err := errors.New("passwords do not match")
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()
	result := db.Where("email = ?", registerRequest.Email).First(&models.User{})

	if ok, status := shouldRegister(result); !ok {
		err := errors.New("user already registered")
		glog.Error(err)
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}

	passwordHash, err := utils.HashPassword(registerRequest.Password)
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Name:         registerRequest.Name,
		Email:        registerRequest.Email,
		PasswordHash: passwordHash,
	}

	result = db.Create(&user)
	if result.Error != nil {
		glog.Error(result.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	jwtToken, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":   jwtToken,
		"user_id": fmt.Sprintf("%d", user.ID),
		"name":    user.Name,
		"email":   user.Email,
	})
}

func Login(ctx *gin.Context) {
	var loginRequest models.LoginRequest

	if err := ctx.BindJSON(&loginRequest); err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if loginRequest.Email == "" || loginRequest.Password == "" {
		err := errors.New("missing required fields")
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	db := database.GetDB()
	result := db.Where("email = ?", loginRequest.Email).First(&user)

	if ok, status := shouldLogin(result); !ok {
		glog.Error(result.Error)
		ctx.JSON(status, gin.H{"error": result.Error.Error()})
		return
	}

	if !utils.CheckPasswordHash(loginRequest.Password, user.PasswordHash) {
		err := errors.New("invalid password")
		glog.Error(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	jwtToken, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":   jwtToken,
		"user_id": fmt.Sprintf("%d", user.ID),
		"name":    user.Name,
		"email":   user.Email,
	})
}

func GetProfile(ctx *gin.Context) {
	userIDStr, exists := ctx.Get("userID")
	if !exists {
		err := errors.New("user not authenticated")
		glog.Error(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	userID, err := strconv.ParseUint(fmt.Sprintf("%v", userIDStr), 10, 32)
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	db := database.GetDB()
	result := db.First(&user, userID)

	if result.Error != nil {
		glog.Error(result.Error)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, user)
}
