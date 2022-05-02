package main

import "fmt"

type JsonElement struct {
	Name  string
	Value string
}

func (jsonElement JsonElement) Append() {
	jsonDoc += fmt.Sprintf("\n\"%s\": \"%s\"", jsonElement.Name, jsonElement.Value)
}
