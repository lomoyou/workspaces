package app

import (
	"github.com/astaxie/beego/validation"
	"go_blog/log"
)

func MarkEorros(errors []*validation.Error) {
	for _, err := range errors {
		log.Infof("err.Key:%v, err.Message:%v", err.Key, err.Message)
	}

}