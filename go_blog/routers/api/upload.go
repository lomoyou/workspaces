package api

import (
	"github.com/gin-gonic/gin"
	"go_blog/log"
	"go_blog/pkg/app"
	"go_blog/pkg/e"
	"go_blog/pkg/upload"
	"net/http"
)

func UploadImage( c *gin.Context) {
	appG := app.Gin{C:c}
	file, image, err :=c.Request.FormFile("image")
	if err != nil {
		log.Warnf("c.Request.FormFile warn, err:%v", err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if image == nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}

	err = upload.CheckImage(fullPath)
	if err != nil {
		log.Warnf("upload.CheckImage warn,err:%v", err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
	}

	if err := c.SaveUploadedFile(image, src); err != nil {
		log.Warnf("c.SaveUploadedFile warn err:%v", err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"image_url": upload.GetImageFullUrl(imageName),
		"image_save_url" : savePath+imageName,
	})
}