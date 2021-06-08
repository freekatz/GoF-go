package template

import "fmt"

type DisplayerTemplate struct {
	displayer Displayer
}

func NewDisplayerTemplate(displayer Displayer) *DisplayerTemplate {
	return &DisplayerTemplate{
		displayer: displayer,
	}
}

func (template *DisplayerTemplate) Display() {
	err := template.displayer.Open()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = template.displayer.Print()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = template.displayer.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
}
