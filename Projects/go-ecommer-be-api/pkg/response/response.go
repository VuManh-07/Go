package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(c *gin.Context, code int, data any) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg[code],
		Data: data,
	})
}

func Error(c *gin.Context, code int) {
	c.JSON(http.StatusBadRequest, Response{
		Code: code,
		Msg:  msg[code],
		Data: nil,
	})
}
