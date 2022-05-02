package main

import (
	"fmt"
)

type JsonElement interface {
	Append()
}

type String struct {
	Value string
}

func (s String) Append() {
	jsonDoc += fmt.Sprintf("\"%s\"", s.Value)
}
