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

type Boollean struct {
	Value bool
}

func (b Boollean) Append(indentationLevel int) {
	jsonDoc += fmt.Sprintf("\"%t\"", b.Value)
}

type Number struct {
	Value int
}

func (i Number) Append(indentationLevel int) {
	jsonDoc += fmt.Sprintf("%v", i.Value)
}

type Null struct {
}

func (n Null) Append(indentationLevel int) {
	jsonDoc += fmt.Sprintf("null")
}

func addTabs(level int) string {
	tabs := ""
	for i := 0; i < level; i++ {
		tabs += "\t"
	}
	return tabs
}
