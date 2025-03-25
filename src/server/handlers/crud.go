package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"webapp/server/database"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"gorm.io/gorm"
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

	CreateFunc    gin.HandlerFunc
	GetFunc       gin.HandlerFunc
	GetAllFunc    gin.HandlerFunc
	UpdateFunc    gin.HandlerFunc
	DeleteFunc    gin.HandlerFunc
	DeleteAllFunc gin.HandlerFunc
}

func getHandler(override, fallback gin.HandlerFunc) gin.HandlerFunc {
	if override != nil {
		return override
	}
	return fallback
}

func RegisterCRUDHandlers[T any](cfg HandlerConfig) {
	basePath := cfg.BasePath
	idPath := fmt.Sprintf("%s/:id", basePath)

	cfg.RouterGroup.POST(basePath, getHandler(cfg.CreateFunc, handleCreate[T]))
	cfg.RouterGroup.GET(basePath, getHandler(cfg.GetAllFunc, handleGetAll[T]))
	cfg.RouterGroup.DELETE(basePath, getHandler(cfg.DeleteAllFunc, handleDeleteAll[T]))

	cfg.RouterGroup.GET(idPath, getHandler(cfg.GetFunc, handleGet[T]))
	cfg.RouterGroup.PUT(idPath, getHandler(cfg.UpdateFunc, handleUpdate[T]))
	cfg.RouterGroup.DELETE(idPath, getHandler(cfg.DeleteFunc, handleDelete[T]))
}

func handleCreate[T any](ctx *gin.Context) {
	var object T

	err := ctx.BindJSON(&object)
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()
	result := db.Create(&object)
	if result.Error != nil {
		glog.Error(result.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, object)
}

func handleGet[T any](ctx *gin.Context) {
	var object T

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()
	result := db.First(&object, id)
	if result.Error != nil {
		glog.Error(result.Error)
		ctx.JSON(getHttpStatusCode(result.Error), gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, object)
}

func handleGetAll[T any](ctx *gin.Context) {
	var objects []T

	db := database.GetDB()
	result := db.Find(&objects)
	if result.Error != nil {
		glog.Error(result.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, objects)
}

func handleUpdate[T any](ctx *gin.Context) {
	var object T

	err := ctx.BindJSON(&object)
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()
	result := db.Save(&object)
	if result.Error != nil {
		glog.Error(result.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, object)
}

func handleDelete[T any](ctx *gin.Context) {
	var object T

	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		glog.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	glog.Error(idStr)

	db := database.GetDB()
	result := db.Delete(&object, id)
	if result.Error != nil {
		glog.Error(result.Error)
		ctx.JSON(getHttpStatusCode(result.Error), gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"rows_affected": result.RowsAffected})
}

func handleDeleteAll[T any](ctx *gin.Context) {
	var object T

	db := database.GetDB()
	result := db.Where("1 = 1").Delete(&object)
	if result.Error != nil {
		glog.Error(result.Error)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"rows_affected": result.RowsAffected})
}
