package routers

import (
	"github.com/KarmaBeLike/doodocs_days/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	router := gin.Default()

	router.POST("/archive/info", handlers.GetArchiveInfoHandler)
	router.POST("/archive/files")

	return router
}
