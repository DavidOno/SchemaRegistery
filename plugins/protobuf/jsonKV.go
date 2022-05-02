package main

import "fmt"

type JsonKV struct {
	Name  string
	Value JsonElement
}

func (jsonElement JsonKV) Append() {
	jsonDoc += fmt.Sprintf("\n\"%s\": ", jsonElement.Name)
	jsonElement.Value.Append()
}
