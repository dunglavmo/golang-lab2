package routes

import (
	"social_app/controllers"
	"social_app/middleware"

	"github.com/gin-gonic/gin"
)

type AlbumRouteController struct {
	AlbumController controllers.AlbumController
}

func NewRouteAlbumController(AlbumController controllers.AlbumController) AlbumRouteController {
	return AlbumRouteController{AlbumController}
}

func (pc *AlbumRouteController) AlbumRoute(rg *gin.RouterGroup) {

	router := rg.Group("albums")
	router.Use(middleware.DeserializeAlbum())
	router.POST("/", pc.AlbumController.CreateAlbum)
	router.GET("/", pc.AlbumController.FindAlbums)
	router.PUT("/:albumId", pc.AlbumController.UpdateAlbum)
	router.GET("/:albumId", pc.AlbumController.FindAlbumById)
	router.DELETE("/:albumId", pc.AlbumController.DeleteAlbum)
}
