package main

import (
	"fmt"
)

type JsonElement interface {
	Append(indentationLevel int)
}

type String struct {
	Value string
}

func (s String) Append(indentationLevel int) {
	jsonDoc += fmt.Sprintf("\"%s\"", s.Value)
}

func addTabs(level int) string {
	tabs := ""
	for i := 0; i < level; i++ {
		tabs += "\t"
	}
	return tabs
}
