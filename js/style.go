package js

import "syscall/js"

type Style struct {
	js.Value
}

func (s Style) Set(name, value string) {
	s.Call("setProperty", name, value)
}