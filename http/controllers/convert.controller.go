package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/navisot/go-exchanger/http/requests"
	"github.com/navisot/go-exchanger/services/interfaces"
	"net/http"
)

// ConvertController has a convertService interface ready to use
type ConvertController struct {
	I interfaces.ConvertServiceInterface
}

// Convert makes the conversion
func (c *ConvertController) Convert(ctx *gin.Context) {

	var data *requests.ConvertRequest

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	result, err := c.I.Convert(data)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})
	return
}
