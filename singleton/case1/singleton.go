package singleton_case1

import (
	"time"
)

type Singleton struct {
	CreateDate string
}

var singleton *Singleton

// 线程不安全的
func GetInstance() *Singleton {
	if singleton == nil {
		singleton = newInstance()
	}
	return singleton
}

// 必须小写保证不可导出
func newInstance() *Singleton {
	return &Singleton{
		CreateDate: time.Now().Format("2006/1/2 15:04:05"),
	}
}
