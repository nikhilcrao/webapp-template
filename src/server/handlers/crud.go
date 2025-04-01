package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"webapp/server/database"
	"webapp/server/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func getHttpStatusCode(err error) int {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound
	} else {
		return http.StatusInternalServerError
	}
}

type HandlerConfig struct {
	BasePath    string
	RouterGroup *gin.RouterGroup
	UserScoped  bool

	CreateFunc    gin.HandlerFunc
	GetFunc       gin.HandlerFunc
	GetAllFunc    gin.HandlerFunc
	UpdateFunc    gin.HandlerFunc
	DeleteFunc    gin.HandlerFunc
	DeleteAllFunc gin.HandlerFunc
}

func getHandler(override, fallback gin.HandlerFunc, userScoped bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("userScoped", userScoped)
		if override != nil {
			override(ctx)
		} else {
			fallback(ctx)
		}
	}
}

func RegisterCRUDHandlers[T any](cfg HandlerConfig) {
	basePath := cfg.BasePath
	idPath := fmt.Sprintf("%s/:id", basePath)

	cfg.RouterGroup.POST(basePath, getHandler(cfg.CreateFunc, handleCreate[T], cfg.UserScoped))
	cfg.RouterGroup.GET(basePath, getHandler(cfg.GetAllFunc, handleGetAll[T], cfg.UserScoped))
	cfg.RouterGroup.DELETE(basePath, getHandler(cfg.DeleteAllFunc, handleDeleteAll[T], cfg.UserScoped))

	cfg.RouterGroup.GET(idPath, getHandler(cfg.GetFunc, handleGet[T], cfg.UserScoped))
	cfg.RouterGroup.PUT(idPath, getHandler(cfg.UpdateFunc, handleUpdate[T], cfg.UserScoped))
	cfg.RouterGroup.DELETE(idPath, getHandler(cfg.DeleteFunc, handleDelete[T], cfg.UserScoped))
}

func handleCreate[T any](ctx *gin.Context) {
	var object T

	err := ctx.BindJSON(&object)
	if err != nil {
		glog.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()
	result := db.Create(&object)
	if result.Error != nil {
		glog.Error(result.Error)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, object)
}

func getQueryBuilder[T any](ctx *gin.Context) *gorm.DB {
	db := database.GetDB()

	var object T
	builder := db.Model(&object).Debug()

	userScoped, exists := ctx.Get("userScoped")
	if !exists {
		glog.Warning("userScoped context not set")
		userScoped = false
	}

	userID, exists := middlewares.GetUserIdFromContext(ctx)
	if !exists {
		glog.Warning("userID not in context")
		userScoped = false
	}

	if userScoped == true {
		builder = builder.Where("user = ?", userID)
	} else {
		builder = builder.Where("1 = 1")
	}

	return builder
}

func handleGet[T any](ctx *gin.Context) {
	var object T

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		glog.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := getQueryBuilder[T](ctx).Preload(clause.Associations).First(&object, id)
	if result.Error != nil {
		glog.Error(result.Error)
		ctx.IndentedJSON(getHttpStatusCode(result.Error), gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, object)
}

func handleGetAll[T any](ctx *gin.Context) {
	var objects []T

	result := getQueryBuilder[T](ctx).Preload(clause.Associations).Find(&objects)
	if result.Error != nil {
		glog.Error(result.Error)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, objects)
}

func handleUpdate[T any](ctx *gin.Context) {
	var object T

	err := ctx.BindJSON(&object)
	if err != nil {
		glog.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()
	result := db.Save(&object)
	if result.Error != nil {
		glog.Error(result.Error)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, object)
}

func handleDelete[T any](ctx *gin.Context) {
	var object T

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		glog.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	glog.Error(idStr)

	result := getQueryBuilder[T](ctx).Delete(&object, id)
	if result.Error != nil {
		glog.Error(result.Error)
		ctx.IndentedJSON(getHttpStatusCode(result.Error), gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"rows_affected": result.RowsAffected})
}

func handleDeleteAll[T any](ctx *gin.Context) {
	var object T

	result := getQueryBuilder[T](ctx).Delete(&object)
	if result.Error != nil {
		glog.Error(result.Error)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"rows_affected": result.RowsAffected})
}
