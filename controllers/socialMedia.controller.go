package controllers

import (
	"net/http"
	"time"

	"final-project-golang/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SocialMediaController struct {
	DB *gorm.DB
}

func NewSocialMediaController(DB *gorm.DB) SocialMediaController {
	return SocialMediaController{DB}
}

func (smc *SocialMediaController) CreateSocialMedia(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload models.CreateSocialMediaRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	now := time.Now()
	newSocialMedia := models.SocialMedia{
		Name:           payload.Name,
		SocialMediaURL: payload.SocialMediaURL,
		UserID:         currentUser.ID,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	result := smc.DB.Create(&newSocialMedia)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":               newSocialMedia.ID,
		"name":             newSocialMedia.Name,
		"social_media_url": newSocialMedia.SocialMediaURL,
		"user_id":          newSocialMedia.UserID,
	})
}

func (smc *SocialMediaController) GetSocialMedias(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var socialMedias []models.SocialMedia
	result := smc.DB.Where("user_id = ?", currentUser.ID).Find(&socialMedias)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unexpected error"})
		return
	}

	// Membuat slice untuk menyimpan respons yang sesuai dengan spesifikasi OpenAPI
	var responseData []gin.H
	for _, socialMedia := range socialMedias {
		responseData = append(responseData, gin.H{
			"id":               socialMedia.ID,
			"name":             socialMedia.Name,
			"social_media_url": socialMedia.SocialMediaURL,
			"user_id":          socialMedia.UserID,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": responseData})
}

func (smc *SocialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	socialMediaID := ctx.Param("socialMediaId")

	var payload models.UpdateSocialMediaRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var updatedSocialMedia models.SocialMedia
	result := smc.DB.First(&updatedSocialMedia, "id = ?", socialMediaID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No social media with that ID exists"})
		return
	}

	updatedSocialMedia.Name = payload.Name
	updatedSocialMedia.SocialMediaURL = payload.SocialMediaURL
	smc.DB.Save(&updatedSocialMedia)

	// Membuat objek JSON yang sesuai dengan spesifikasi OpenAPI
	responseData := gin.H{
		"id":               updatedSocialMedia.ID,
		"name":             updatedSocialMedia.Name,
		"social_media_url": updatedSocialMedia.SocialMediaURL,
		"user_id":          updatedSocialMedia.UserID,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": responseData})
}

func (smc *SocialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	socialMediaID := ctx.Param("socialMediaId")

	result := smc.DB.Delete(&models.SocialMedia{}, "id = ?", socialMediaID)
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No social media with that ID exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
