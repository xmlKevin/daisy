package handler

import (
	jwt "daisy/pkg/jwtauth"
	"daisy/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NoFound(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	logger.Infof("NoRoute claims: %#v\n", claims)
	c.JSON(http.StatusOK, gin.H{
		"code":    "404",
		"message": "not found",
	})
}
