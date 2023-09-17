package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"social_app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AlbumController struct {
	DB *gorm.DB
}

func NewAlbumController(DB *gorm.DB) AlbumController {
	return AlbumController{DB}
}

func (ac *AlbumController) CreateAlbum(ctx *gin.Context) {
	var payload *models.CreateAlbumRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newAlbum := models.Album{
		Name:        payload.Name,
		Description: payload.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	result := ac.DB.Create(&newAlbum)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Album with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newAlbum})
}

func (ac *AlbumController) UpdateAlbum(ctx *gin.Context) {
	albumId := ctx.Param("albumId")

	var payload *models.UpdateAlbum
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedAlbum models.Album
	result := ac.DB.First(&updatedAlbum, "id = ?", albumId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No Album with that title exists"})
		return
	}
	now := time.Now()
	AlbumToUpdate := models.Album{
		Name:        payload.Name,
		Description: payload.Description,
		CreatedAt:   updatedAlbum.CreatedAt,
		UpdatedAt:   now,
	}

	ac.DB.Model(&updatedAlbum).Updates(AlbumToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedAlbum})
}

func (ac *AlbumController) FindAlbumById(ctx *gin.Context) {
	albumId := ctx.Param("albumId")

	var Album models.Album
	result := ac.DB.First(&Album, "id = ?", albumId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No Album with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": Album})
}

func (ac *AlbumController) FindAlbums(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var Albums []models.Album
	results := ac.DB.Limit(intLimit).Offset(offset).Find(&Albums)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(Albums), "data": Albums})
}

func (ac *AlbumController) DeleteAlbum(ctx *gin.Context) {
	albumId := ctx.Param("albumId")

	result := ac.DB.Delete(&models.Album{}, "id = ?", albumId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No Album with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})

}
