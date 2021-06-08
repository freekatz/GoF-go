package template

type Displayer interface {
	Open() error // 方法的大小写 (可导出性) 根据不同情形来决定
	Print() error
	Close() error
}
