package middleware

import (
	"net/http"

	"social_app/config"
	"social_app/models"

	"github.com/gin-gonic/gin"
)

func DeserializeAlbum() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// var token string
		// cookie, err := ctx.Cookie("token")

		// authorizationHeader := ctx.Request.Header.Get("Authorization")
		// fields := strings.Fields(authorizationHeader)

		// if len(fields) != 0 && fields[0] == "Bearer" {
		// 	token = fields[1]
		// } else if err == nil {
		// 	token = cookie
		// }

		// if token == "" {
		// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
		// 	return
		// }

		// config.LoadEnv()
		// sub, err := utils.ValidateToken(token, os.Getenv("TOKEN_SECRET"))
		// if err != nil {
		// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
		// 	return
		// }

		var album models.Album
		result := config.DB.First(&album, "id = ?")
		if result.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the album belonging to this token no logger exists"})
			return
		}

		ctx.Set("currentAlbum", album)
		ctx.Next()
	}
}
