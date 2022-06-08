/*
  @Author : xmlKevin
*/

package process

import (
	"daisy/apis/process"
	"daisy/middleware"

	//"daisy/apis/process"
	//"daisy/middleware"
	jwt "daisy/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

func RegisterTplRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	tplRouter := v1.Group("/tpl").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		tplRouter.GET("", process.TemplateList)
		tplRouter.POST("", process.CreateTemplate)
		tplRouter.PUT("", process.UpdateTemplate)
		tplRouter.DELETE("", process.DeleteTemplate)
		tplRouter.GET("/details", process.TemplateDetails)
		tplRouter.POST("/clone/:id", process.CloneTemplate)
	}
}
