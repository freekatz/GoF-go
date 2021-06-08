package singleton_case3

import (
	"sync"
	"time"
)

type Singleton struct {
	CreateDate string
}

var singleton *Singleton
var once sync.Once

// 线程安全的
func GetInstance() *Singleton {
	once.Do(func() {
		if singleton == nil {
			singleton = newInstance()
		}
	})
	return singleton
}

// 必须小写保证不可导出
func newInstance() *Singleton {
	return &Singleton{
		CreateDate: time.Now().Format("2006/1/2 15:04:05"),
	}
}
