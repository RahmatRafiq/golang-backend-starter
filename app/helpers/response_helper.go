package helpers

import (
	"github.com/gin-gonic/gin"
)

// Data
type Params struct {
	Data    *[]interface{} `json:"data"`
	Item    *interface{}  `json:"item"`
	Message *string 		 `json:"message"`
	Token   *string 		 `json:"token"`
}

// OK Response
func OK(ctx *gin.Context, params *Params) {
	ctx.JSON(200, map[string]interface{}{
		"status":  "OK",
		"data":    &params.Data,
		"item":    &params.Item,
		"message": &params.Message,
		"token":   &params.Token,
	})
}
