package iterator

import "fmt"

// ===
// 定义测试 Object
// ===

type Object struct {
	ID int
}

type SecretType int

const (
	Love SecretType = iota
	Life
	NoneSecret
)

func NewObject(id int) *Object {
	return &Object{
		ID: id,
	}
}

func (o *Object) Secret() SecretType {
	switch {
	case o.ID%3 == 0:
		return Love
	case o.ID%3 == 1:
		return Life
	default:
		return NoneSecret
	}
}

func (t SecretType) String() string {
	switch t {
	case Love:
		return "Love Secret"
	case Life:
		return "Life Secret"
	default:
		return "None Secret"
	}
}

func (o *Object) String() string {
	return fmt.Sprintf("Object with ID: %d, has secert called: %s.\n", o.ID, o.Secret().String())
}
