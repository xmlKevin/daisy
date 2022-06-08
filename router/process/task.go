package process

import (
	"daisy/apis/process"
	"daisy/middleware"
	jwt "daisy/pkg/jwtauth"

	"github.com/gin-gonic/gin"
)

/*
  @Author : xmlKevin
*/

func RegisterTaskRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	taskRouter := v1.Group("/task").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		taskRouter.GET("", process.TaskList)
		taskRouter.GET("/details", process.TaskDetails)
		taskRouter.POST("", process.CreateTask)
		taskRouter.PUT("", process.UpdateTask)
		taskRouter.DELETE("", process.DeleteTask)
	}
}
