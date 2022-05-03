package main

import "fmt"

type JsonKV struct {
	Name  string
	Value JsonElement
}

func (jsonElement JsonKV) Append(indentationLevel int) {
	jsonDoc += fmt.Sprintf("\n%s\"%s\": ", addTabs(indentationLevel), jsonElement.Name)
	jsonElement.Value.Append(indentationLevel)
}
