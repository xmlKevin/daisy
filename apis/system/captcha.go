package system

import (
	"daisy/tools/app"
	"daisy/tools/captcha"
	"fmt"

	"github.com/gin-gonic/gin"
)

/*
  @Author : xmlKevin
*/

func GenerateCaptchaHandler(c *gin.Context) {
	id, b64s, err := captcha.DriverDigitFunc()
	if err != nil {
		app.Error(c, -1, err, fmt.Sprintf("验证码获取失败, %v", err.Error()))
		return
	}
	app.Custum(c, gin.H{
		"code": 200,
		"data": b64s,
		"id":   id,
		"msg":  "success",
	})
}
