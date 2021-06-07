package template

type Displayer interface {
	Open() error
	Print() error
	Close() error
}
