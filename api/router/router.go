package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	//Define all endpoints
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Chill")
	})

	return router
}
