package main

type JsonArray struct {
	Objects []JsonElement
}

func (jsonList JsonArray) Append(indentationLevel int) {
	jsonDoc += "["
	for index, element := range jsonList.Objects {
		if index > 0 {
			jsonDoc += ","
		}
		element.Append(indentationLevel)
	}
	jsonDoc += "]"
}
