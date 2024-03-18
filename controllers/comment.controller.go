package controllers

import (
	"net/http"
	"time"

	"final-project-golang/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentController struct {
	DB *gorm.DB
}

func NewCommentController(DB *gorm.DB) CommentController {
	return CommentController{DB}
}

func (cc *CommentController) CreateComment(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreateCommentRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Validasi field message
	if payload.Message == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Message is required"})
		return
	}

	now := time.Now()
	newComment := models.Comment{
		Message:   payload.Message,
		PhotoID:   payload.PhotoID,
		UserID:    currentUser.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := cc.DB.Create(&newComment)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":       newComment.ID,
		"message":  newComment.Message,
		"photo_id": newComment.PhotoID,
		"user_id":  newComment.UserID,
	})
}

func (cc *CommentController) GetComments(ctx *gin.Context) {
	var comments []models.Comment
	result := cc.DB.Find(&comments)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unexpected error"})
		return
	}

	// Membuat slice untuk menyimpan respons yang sesuai dengan spesifikasi OpenAPI
	var responseData []gin.H
	for _, comment := range comments {
		responseData = append(responseData, gin.H{
			"id":       comment.ID,
			"message":  comment.Message,
			"photo_id": comment.PhotoID,
			"user_id":  comment.UserID,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": responseData})
}

func (cc *CommentController) GetCommentByID(ctx *gin.Context) {
	commentID := ctx.Param("commentId")

	var comment models.Comment
	result := cc.DB.First(&comment, "id = ?", commentID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No comment with that ID exists"})
		return
	}

	// Membuat objek JSON yang sesuai dengan spesifikasi OpenAPI
	responseData := gin.H{
		"id":       comment.ID,
		"message":  comment.Message,
		"photo_id": comment.PhotoID,
		"user_id":  comment.UserID,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": responseData})
}

func (cc *CommentController) UpdateComment(ctx *gin.Context) {
	commentID := ctx.Param("commentId")

	var payload models.UpdateCommentRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Validasi field message
	if payload.Message == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Message is required"})
		return
	}

	var updatedComment models.Comment
	result := cc.DB.First(&updatedComment, "id = ?", commentID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No comment with that ID exists"})
		return
	}

	updatedComment.Message = payload.Message
	cc.DB.Save(&updatedComment)

	// Membuat objek JSON yang sesuai dengan spesifikasi OpenAPI
	responseData := gin.H{
		"id":       updatedComment.ID,
		"message":  updatedComment.Message,
		"photo_id": updatedComment.PhotoID,
		"user_id":  updatedComment.UserID,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": responseData})
}

func (cc *CommentController) DeleteComment(ctx *gin.Context) {
	commentID := ctx.Param("commentId")

	result := cc.DB.Delete(&models.Comment{}, "id = ?", commentID)
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No comment with that ID exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
