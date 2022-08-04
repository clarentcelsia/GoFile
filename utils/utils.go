package utils

import (
	"net/http"
	m "pdf/models"

	"github.com/gin-gonic/gin"
)

// Hit rest api

func HandleResponse(c *gin.Context, result interface{}, message string, status int) {
	var resp m.PageResponse
	resp = m.PageResponse{
		Message: message,
		Result:  result,
		Status:  status,
	}
	c.JSON(http.StatusOK, resp)
}
