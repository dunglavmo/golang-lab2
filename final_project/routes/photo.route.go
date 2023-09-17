package routes

import (
	"social_app/controllers"
	"social_app/middleware"

	"github.com/gin-gonic/gin"
)

type PhotoRouteController struct {
	PhotoController controllers.PhotoController
}

func NewRoutePhotoController(PhotoController controllers.PhotoController) PhotoRouteController {
	return PhotoRouteController{PhotoController}
}

func (pc *PhotoRouteController) PhotoRoute(rg *gin.RouterGroup) {

	router := rg.Group("photos")
	router.Use(middleware.DeserializeUser())
	router.Use(middleware.DeserializeAlbum())
	router.POST("/", pc.PhotoController.CreatePhoto)
	router.GET("/", pc.PhotoController.FindPhotos)
	router.PUT("/:photoId", pc.PhotoController.UpdatePhoto)
	router.GET("/:photoId", pc.PhotoController.FindPhotoById)
	router.DELETE("/:photoId", pc.PhotoController.DeletePhoto)
}
