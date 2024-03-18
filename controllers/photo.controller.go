package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"final-project-golang/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhotoController struct {
	DB *gorm.DB
}

func NewPhotoController(DB *gorm.DB) PhotoController {
	return PhotoController{DB}
}

func (pc *PhotoController) CreatePhoto(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreatePhotoRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	now := time.Now()
	newPhoto := models.Photo{
		Title:     payload.Title,
		Caption:   payload.Caption,
		PhotoURL:  payload.PhotoURL,
		UserID:    currentUser.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := pc.DB.Create(&newPhoto)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"message": "Photo with that title already exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":        newPhoto.ID,
		"caption":   newPhoto.Caption,
		"title":     newPhoto.Title,
		"photo_url": newPhoto.PhotoURL,
		"user_id":   newPhoto.UserID,
	})
}

func (pc *PhotoController) UpdatePhoto(ctx *gin.Context) {
	photoID := ctx.Param("photoId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdatePhoto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedPhoto models.Photo
	result := pc.DB.First(&updatedPhoto, "id = ?", photoID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No photo with that ID exists"})
		return
	}
	now := time.Now()
	photoToUpdate := models.Photo{
		Title:     payload.Title,
		Caption:   payload.Caption,
		PhotoURL:  payload.PhotoURL,
		UserID:    currentUser.ID,
		CreatedAt: updatedPhoto.CreatedAt,
		UpdatedAt: now,
	}

	pc.DB.Model(&updatedPhoto).Updates(photoToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedPhoto})
}

func (pc *PhotoController) FindPhotoByID(ctx *gin.Context) {
	photoID := ctx.Param("photoId")

	var photo models.Photo
	result := pc.DB.First(&photo, "id = ?", photoID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No photo with that ID exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": photo})
}

func (pc *PhotoController) FindPhotos(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var photos []models.Photo
	results := pc.DB.Limit(intLimit).Offset(offset).Find(&photos)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(photos), "data": photos})
}

func (pc *PhotoController) DeletePhoto(ctx *gin.Context) {
	photoID := ctx.Param("photoId")

	result := pc.DB.Delete(&models.Photo{}, "id = ?", photoID)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No photo with that ID exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
