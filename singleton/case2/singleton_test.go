package singleton_case2

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	concurrencyTest() /// 会输出 false 结果
	commonTest()
}

func commonTest() {
	singleton1 := GetInstance()
	time.Sleep(3 * time.Second)
	singleton2 := GetInstance()

	fmt.Println(singleton1.CreateDate, singleton2.CreateDate, singleton1 == singleton2)
}

func concurrencyTest() {

	var singleton1 *Singleton
	var singleton2 *Singleton

	go func() {
		singleton1 = GetInstance()
	}()

	func() {
		singleton2 = GetInstance()
	}()

	time.Sleep(3 * time.Second)

	fmt.Println(singleton1, singleton2, singleton1 == singleton2)

	singleton3 := GetInstance()

	fmt.Println(singleton1.CreateDate, singleton3.CreateDate, singleton1 == singleton3, singleton2 == singleton3)

}
