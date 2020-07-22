package upload

import (
	"fmt"
	"go_blog/file"
	"go_blog/log"
	"go_blog/util"
	"go_blog/pkg/setting"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

//GetImageFullUrl get the full access path
func GetImageFullUrl (name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetImageFullPath() + name
}
//GetImageName get image name
func GetImageName (name string) string{
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

//GetIamgePath get save path
func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

//GetImageFullPath get full save path
func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

//
func CheckImageExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}
	return false
}

//
func CheckImageSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Errorf("CheckImageSize err: %v", err)
		return false
	}

	return size <= setting.AppSetting.ImageMaxSize
}

//
func CheckImage(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		log.Errorf("os.Getwd err:%v", err)
	}

	err = file.IsNotExistMKDir(dir + "/" +src)
	if err != nil {
		log.Errorf("file.IsNotExistMKDir err:%v", err)
	}

	perm := file.CheckPermission(src)
	if perm {
		return fmt.Errorf("file.Checkpermis denied src: %s", src)
	}
	return nil
}
