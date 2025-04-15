package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	appVersion "template-api-go/internal/function"
)

func GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": appVersion.GetAppVersion(),
	})
}
