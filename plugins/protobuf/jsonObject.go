package main

import "fmt"

type JsonObject struct {
	Elements []JsonElement
}

func (jsonObject JsonObject) Append(indentationLevel int) {
	jsonDoc += fmt.Sprintf("{")
	for index, element := range jsonObject.Elements {
		if index > 0 {
			jsonDoc += ","
		}
		element.Append(indentationLevel + 1)
	}
	jsonDoc += fmt.Sprintf("\n%s}", addTabs(indentationLevel))
}
