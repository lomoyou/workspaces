package util

import "go_blog/pkg/setting"

func Setup() {
	jwtSercet = []byte(setting.AppSetting.JwtSecret)
}
