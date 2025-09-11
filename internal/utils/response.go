package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, message string, data interface{}) {
    c.JSON(http.StatusOK, APIResponse{
        Status:  "success",
        Message: message,
        Data:    data,
    })
}

func Created(c *gin.Context, message string, data interface{}) {
    c.JSON(http.StatusCreated, APIResponse{
        Status:  "success",
        Message: message,
        Data:    data,
    })
}

func Error(c *gin.Context, code int, message string) {
    c.JSON(code, APIResponse{
        Status:  "error",
        Message: message,
    })
}
