package util

import (
	"github.com/unknwon/com"
	"github.com/gin-gonic/gin"
	"go_blog/pkg/setting"
)

func GetPage( c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page >0 {
		result = (page -1) * setting.AppSetting.PageSize
	}

	return  result
}