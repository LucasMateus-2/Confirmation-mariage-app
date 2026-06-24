package response

import "github.com/gin-gonic/gin"

type Body struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func OK(c *gin.Context, data any) {
	c.JSON(200, Body{Data: data})
}

func Created(c *gin.Context, data any) {
	c.JSON(201, Body{Data: data})
}

func BadRequest(c *gin.Context, msg string) {
	c.JSON(400, Body{Error: msg})
}

func Unauthorized(c *gin.Context) {
	c.JSON(401, Body{Error: "não autorizado"})
}

func NotFound(c *gin.Context, msg string) {
	c.JSON(404, Body{Error: msg})
}

func InternalError(c *gin.Context) {
	c.JSON(500, Body{Error: "erro interno no servidor"})
}
