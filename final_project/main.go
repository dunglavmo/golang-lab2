// package main

// import (
// 	"log"
// 	"net/http"
// 	"os"

// 	"social_app/config"

// 	"github.com/gin-gonic/gin"
// )

// var (
// 	server *gin.Engine
// )

// func init() {
// 	config.LoadEnv()

// 	config.ConnectDB()

// 	server = gin.Default()
// }

// func main() {
// 	router := server.Group("/api")
// 	router.GET("/healthchecker", func(ctx *gin.Context) {
// 		message := "Welcome to Golang "
// 		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
// 	})

// 	log.Fatal(server.Run(":" + os.Getenv("PORT")))
// }

package main

import (
	"log"
	"net/http"
	"os"

	"social_app/config"
	"social_app/controllers"
	"social_app/routes"

	"github.com/gin-gonic/gin"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController
)

func init() {
	config.LoadEnv()

	config.ConnectDB()

	AuthController = controllers.NewAuthController(config.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(config.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	server = gin.Default()
}

func main() {
	config.LoadEnv()

	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{"http://localhost:8000", os.Getenv("ClientOrigin")}
	// corsConfig.AllowCredentials = true

	// server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	log.Fatal(server.Run(":" + os.Getenv("PORT")))
}
