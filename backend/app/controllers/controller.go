package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

// @Summary	Show a hello world message
// @Tags		root
// @Accept		json
// @Produce	json
// @Success	200	{string} string	"Hello World"
// @Router		/ [get]
func (*Controller) HelloWorld(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, "Hello World")
}
