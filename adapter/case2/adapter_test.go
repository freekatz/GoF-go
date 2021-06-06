package adapter_case2

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	object := NewImplementedObject(1)
	objectTest(object)
}

func objectTest(object Object) {
	secret := object.Secret()
	// Info() 可以正常调用, 却不会 panic
	fmt.Println("Secret: ", secret.String(), "\nInfo(): ", object.Info())
}
