package dashboard

import (
	"daisy/apis/dashboard"
	"daisy/middleware"
	jwt "daisy/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

/*
  @Author : lanyulei
*/

func RegisterDashboardRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	classify := v1.Group("/dashboard").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		classify.GET("", dashboard.InitData)
	}
}
