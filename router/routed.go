package router

import (
	modules "api-go/api"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	route := gin.Default()

	route.POST("/books", modules.AddedBook)
	route.GET("/books", modules.GetAllBook)
	route.GET("/books/:ID", modules.GetBookByID)
	route.PUT("/books/:ID", modules.UpdateBook)
	route.DELETE("/books/:ID", modules.DeleteBook)
	return route
}
