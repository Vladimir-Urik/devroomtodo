package routes

import (
	v1 "devroomtodo/routes/v1"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	server := gin.Default()
	v1.RegisterRoutes(server.Group("/v1"))
	return server
}
