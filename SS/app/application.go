// Starta a aplicação junto com o arquivo .main

package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	router.Run(":8080")
}
