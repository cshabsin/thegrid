package attr

import "fmt"

type Attr struct {
	Name  string
	Value interface{}
}

func Make(name string, value interface{}) Attr {
	return Attr{Name: name, Value: value}
}

func Class(value string) Attr {
	return Attr{Name: "class", Value: value}
}

func Translate(x, y float64) Attr {
	return Make("transform", fmt.Sprintf("translate(%f,%f)", x, y))
}
