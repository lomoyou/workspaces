package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go_blog/log"
	"go_blog/pkg/app"
	"go_blog/pkg/e"
	"go_blog/pkg/export"
	"go_blog/pkg/setting"
	"go_blog/service/tag_service"
	"go_blog/util"
	"net/http"

	"github.com/astaxie/beego/validation"
)

func GetTags( c *gin.Context) {
	appG :=app.Gin{C:c}
	name := c.Query("name")
	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	tagService := tag_service.Tag{
		Name : name,
		State: state,
		PageNum: util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}
	tags, err := tagService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, nil)
		return
	}

	count, err := tagService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"list":tags,
		"total":count,
	})
}

type AddTagForm struct {
	Name   string  `form:"name" valid:"Required;Maxsize(100)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
	State     int    `form:"state" valid:"Range(0,1)"`
}

func AddTag(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddTagForm
	)

	httpcode, errcode := app.BindAndValid(c, &form)
	if errcode != e.SUCCESS {
		appG.Response(httpcode, errcode, nil)
		return
	}

	tagService := tag_service.Tag{
		Name: form.Name,
		CreatedBy: form.CreatedBy,
		State: form.State,
	}
	exists, err := tagService.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_TAG, nil)
		return
	}

	err = tagService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditTagFrom struct {
	ID         int    `form:"id" valid:"Required;Min(1)"`
	Name       string `form:"name" valid:"Required;MaxSize(100)"`
	ModifiedBy string `form:"modified_by" valid:"Required;MaxSize(100)"`
	State      int    `form:"state" valid:"Range(0,1)"`
}

func EditTag(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		from = EditTagFrom{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpcode, errcode := app.BindAndValid(c, &from)
	if errcode != e.SUCCESS {
		appG.Response(httpcode, errcode,nil)
		return
	}

	tagService := tag_service.Tag{
		ID: from.ID,
		Name: from.Name,
		ModifiedBy: from.ModifiedBy,
		State: from.State,
	}
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = tagService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK,e.SUCCESS, nil)
}

func DeletTag( c *gin.Context) {
	appG := app.Gin{C:c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkEorros(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil )
	}

	tagService := tag_service.Tag{ID: id}
	exists, err := tagService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	filename, err := tagService.Export()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXPORT_TAG_FAIL, nil )
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"export_url" : export.GetExcelFullUrl(filename),
		"export_save_url" : export.GetExcelPath() + filename,
	})

}

func ExportTag(c *gin.Context) {
	appG := app.Gin{C: c}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		log.Infof("c.Request.FormFile failed,err:%v", err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	tagService := tag_service.Tag{}
	err = tagService.Import(file)
	if err != nil {
		log.Infof("tagService.Import failed,err:%v", err)
		appG.Response(http.StatusInternalServerError, e.ERROR_IMPORT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}