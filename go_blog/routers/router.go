package routers

import (
	"github.com/gin-gonic/gin"
	"go_blog/middleware/jwt"
	"go_blog/pkg/export"
	"go_blog/pkg/qrcode"
	"go_blog/pkg/upload"
	"go_blog/routers/api"
	v1 "go_blog/routers/api/v1"
	"net/http"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLFiles("./template/login.html")
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.GET("/", func (c *gin.Context){
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags", v1.DeletTag)
		//导出标签
		r.POST("/tags/export", v1.ExportTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticle)
		//获取指定文章
		apiv1.GET("articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("articles/:id", v1.DeleteArticle)

		//生成文章海报
		apiv1.POST("articles/poster/generate", v1.GenerateArticlePoster)
	}

	r.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "王光浩傻逼",
		})
	})

	return r
}
