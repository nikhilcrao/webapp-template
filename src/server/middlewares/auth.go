package middlewares

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"strings"
	"webapp/server/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("test1")
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			ctx.Abort()
			return
		}

		log.Println("test2")
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header format must be Bearer {token}"})
			ctx.Abort()
			return
		}

		log.Println("test3")
		claims, err := utils.ValidateJWT(parts[1])
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			ctx.Abort()
			return
		}

		log.Println("test4")
		userID, err := strconv.ParseUint(claims.UserID, 10, 32)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user_id in token"})
			ctx.Abort()
			return
		}

		log.Println("test5")
		ctx.Set("userID", userID)
		ctx.Set("email", claims.Email)

		ctx.Next()
	}
}

func GetUserIdFromContext(ctx context.Context) (uint, bool) {
	userID, ok := ctx.Value("userID").(uint)
	return userID, ok
}
