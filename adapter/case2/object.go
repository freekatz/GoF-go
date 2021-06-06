package adapter_case2

// ===
// Object 接口定义需要实现的方法,
// 其中 `mustEmbedUnimplementedObject()` 之上作为一种约定,
// 用来提醒开发者, 这个接口可以部分实现, 但是必须要混合一个默认的完全实现接口方法的结构体
// ===

type Object interface {
	Secret() SecretType
	Info() string
	mustEmbedUnimplementedObject()
}

type SecretType int

const (
	Love SecretType = iota
	Life
	NoneSecret
)

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
