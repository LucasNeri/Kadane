// Controller onde recebe os parâmetros e seta ele nas variaveis declaradas

package maxsum

import (
	"net/http"

	"SS/domain/maxsum"
	"SS/services"
	"SS/utils/errors/rest_errors"

	"github.com/gin-gonic/gin"
)

var (
	MaxSumController MaxSumInterface = &maxSumController{}
)

type MaxSumInterface interface {
	MaxSum(c *gin.Context)
}

type maxSumController struct {
}

func (cont *maxSumController) MaxSum(c *gin.Context) {
	var n maxsum.Numbers

	if err := c.ShouldBindJSON(&n); err != nil {
		restErr := rest_errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, countErr := services.MaxSumService.MaxSum(n)
	if countErr != nil {
		c.JSON(countErr.Status, countErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
