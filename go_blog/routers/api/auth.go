package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go_blog/log"
	"go_blog/pkg/app"
	"go_blog/pkg/e"
	"go_blog/service/auth_service"
	"go_blog/util"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valiad := validation.Validation{}
	//username := c.Query("username")
	//password := c.Query("password")
	username := c.PostForm("username")
	password := c.PostForm("password")

	a := auth{Username: username, Password: password}
	ok, _ := valiad.Valid(&a)
	if !ok {
		app.MarkEorros(valiad.Errors)
		log.Errorf("valiad.Errors:%v", valiad.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{Username: username, Password: password}
	isexist, err := authService.Check()
	if err != nil {
		log.Errorf("authService.Check failed, err:%v", err)
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isexist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})

}
