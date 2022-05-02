package main

import "fmt"

type JsonList struct {
	Name    string
	Objects []JsonObject
}

func (jsonList JsonList) Append() {
	jsonDoc += fmt.Sprintf("\n\"%s\": [", jsonList.Name)
	for index, element := range jsonList.Objects {
		if index > 0 {
			jsonDoc += ","
		}
		element.Append()
	}
	jsonDoc += "]"
}
