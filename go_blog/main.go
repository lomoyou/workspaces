package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_blog/gredis"
	"go_blog/log"
	"go_blog/models"
	"go_blog/pkg/setting"
	"go_blog/routers"
	"go_blog/util"
	"net/http"
)

func init() {
	setting.Setup()
	log.Setup()
	models.Setup()
	gredis.Setup()
	util.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderbytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderbytes,
	}

	log.Infof("start http server listening %s", endPoint)

	server.ListenAndServe()
}
