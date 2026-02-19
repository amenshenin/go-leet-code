package easy

import "fmt"

type First struct {
	IsComplete int
	Name       string
	Descr      string
}

func (f *First) Run() error {
	fmt.Printf("%s", "Hello World!")
	return nil
}
