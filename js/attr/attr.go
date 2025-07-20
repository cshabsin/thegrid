package attr

import "fmt"

type Attr struct {
	Name  string
	Value any
}

func Make(name string, value any) Attr {
	return Attr{Name: name, Value: value}
}

func Y(value int) Attr {
	return Attr{"y", value}
}

func X(value int) Attr {
	return Attr{"x", value}
}

func Class(value string) Attr {
	return Attr{"class", value}
}

func Draggable(value bool) Attr {
	return Attr{"draggable", value}
}

func Href(value string) Attr {
	return Attr{"href", value}
}

func ID(value string) Attr {
	return Attr{"id", value}
}



func Type(value string) Attr {
	return Attr{"type", value}
}

func Translate(x, y float64) Attr {
	return Make("transform", fmt.Sprintf("translate(%f,%f)", x, y))
}
