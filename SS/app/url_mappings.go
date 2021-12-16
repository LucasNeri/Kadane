// Mapeia as rotas

package app

import (
	"SS/controllers/maxsum"

	"github.com/gin-gonic/gin"
)

func mapRoutesMaxsum(api *gin.RouterGroup) {
	maxSumGroup := api.Group("maxsum")

	maxSumGroup.POST("", maxsum.MaxSumController.MaxSum)
}

func mapUrls() {
	router.Use(CORSMiddleware())

	api := router.Group("")

	api.OPTIONS("/*path", CORSMiddleware())

	mapRoutesMaxsum(api)
}
