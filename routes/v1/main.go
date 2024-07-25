package v1

import "github.com/gin-gonic/gin"

// RegisterRoutes registers all the routes for the application under the provided router group.
// Currently, it only registers the Todo routes.
func RegisterRoutes(router *gin.RouterGroup) {
	registerTodosRoutes(router)
}
