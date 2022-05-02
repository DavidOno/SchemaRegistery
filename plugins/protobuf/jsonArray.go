package main

type JsonArray struct {
	Objects []JsonObject
}

func (jsonList JsonArray) Append() {
	jsonDoc += "\n["
	for index, element := range jsonList.Objects {
		if index > 0 {
			jsonDoc += ","
		}
		element.Append()
	}
	jsonDoc += "\n]"
}
