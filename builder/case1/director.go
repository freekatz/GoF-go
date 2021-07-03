package case1

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (director *Director) Build(title, body string) {
	director.builder.MakeHead()
	director.builder.MakeTitle(title)
	director.builder.MakeBody(body)
	director.builder.MakeFooter()
}
