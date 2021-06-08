package singleton_case2

import (
	"time"
)

type Singleton struct {
	CreateDate string
}

var singleton *Singleton = newInstance()

func GetInstance() *Singleton {
	return singleton
}

// 必须小写保证不可导出
func newInstance() *Singleton {
	return &Singleton{
		CreateDate: time.Now().Format("2006/1/2 15:04:05"),
	}
}
