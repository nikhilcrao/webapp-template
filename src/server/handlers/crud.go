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

func getObjectID(ctx *gin.Context) (uint, error) {
	id64, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		glog.Error(err)
		return 0, err
	}
	return uint(id64), nil
}

func getUserID(ctx *gin.Context) (uint, error) {
	userID, exists := middlewares.GetUserIdFromContext(ctx)
	if !exists {
		return 0, errors.New("UserID not found")
	}
	return userID, nil
}

func getHttpStatusCode(err error) int {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound
	} else {
		return http.StatusInternalServerError
	}
}

type Handler[T any] interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	DeleteAll(ctx *gin.Context)

	Validate(*T) error
	GetQueryBuilder(uint) *gorm.DB
	QueryUserScoped() bool
	QueryPreloadAssociations() bool
}

func RegisterHandlers[T any](basePath string, group *gin.RouterGroup, handler Handler[T]) {
	group.GET(basePath, handler.GetAll)
	group.POST(basePath, handler.Create)
	group.DELETE(basePath, handler.DeleteAll)

	idPath := fmt.Sprintf("%s/:id", basePath)
	group.GET(idPath, handler.Get)
	group.PUT(idPath, handler.Update)
	group.DELETE(idPath, handler.Delete)
}

// BaseHandler implements Handler interface with the base implementation.
//
// Models may provide a separate implementation or provide overrides as needed.
type BaseHandler[T any] struct {
}

func (h BaseHandler[T]) QueryPreloadAssociations() bool {
	return true
}

func (h BaseHandler[T]) QueryUserScoped() bool {
	return true
}

func (h BaseHandler[T]) Validate(object *T) error {
	return nil
}

func (h BaseHandler[T]) GetQueryBuilder(userID uint) *gorm.DB {
	var object T
	db := database.GetDB().Model(&object).Debug()
	if h.QueryUserScoped() && userID != 0 {
		db = db.Where("user_id = ?", userID)
	} else {
		db = db.Where("1 = 1")
	}
	return db
}

func (h BaseHandler[T]) Create(ctx *gin.Context) {
	var object T

	err := ctx.BindJSON(&object)
	if err != nil {
		glog.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Validate(&object); err != nil {
		glog.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	db := h.GetQueryBuilder(0)
	if err := db.Create(&object).Error; err != nil {
		glog.Error(err)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, object)
}

func (h BaseHandler[T]) GetAll(ctx *gin.Context) {
	var objects []T

	userID, err := getUserID(ctx)
	if err != nil {
		glog.Warning(err)
	}

	db := h.GetQueryBuilder(userID)
	err = db.Find(&objects).Error
	if err != nil {
		glog.Error(err)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, objects)
}

func (h BaseHandler[T]) Get(ctx *gin.Context) {
	var object T

	objectID, err := getObjectID(ctx)
	if err != nil {
		glog.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := getUserID(ctx)
	if err != nil {
		glog.Warning(err)
	}

	db := h.GetQueryBuilder(userID)
	if h.QueryPreloadAssociations() {
		db = db.Preload(clause.Associations)
	}

	if err = db.First(&object, objectID).Error; err != nil {
		glog.Error(err)
		ctx.IndentedJSON(getHttpStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, object)
}

func (h BaseHandler[T]) Update(ctx *gin.Context) {
	var object T

	err := ctx.BindJSON(&object)
	if err != nil {
		glog.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := getUserID(ctx)
	if err != nil {
		glog.Warning(err)
	}

	db := h.GetQueryBuilder(userID)
	if err := db.Save(&object).Error; err != nil {
		glog.Error(err)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, object)
}

func (h BaseHandler[T]) Delete(ctx *gin.Context) {
	var object T

	objectID, err := getObjectID(ctx)
	if err != nil {
		glog.Error(err)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := getUserID(ctx)
	if err != nil {
		glog.Warning(err)
	}

	db := h.GetQueryBuilder(userID)
	result := db.Delete(&object, objectID)
	if err := result.Error; err != nil {
		glog.Error(err)
		ctx.IndentedJSON(getHttpStatusCode(err), gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"rows_affected": result.RowsAffected})
}

func (h BaseHandler[T]) DeleteAll(ctx *gin.Context) {
	var object T

	userID, err := getUserID(ctx)
	if err != nil {
		glog.Warning(err)
	}

	db := h.GetQueryBuilder(userID)
	result := db.Delete(&object)
	if err := result.Error; err != nil {
		glog.Error(err)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"rows_affected": result.RowsAffected})
}
