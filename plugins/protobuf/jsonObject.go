package main

type JsonObject struct {
	Elements []JsonElement
}

func (jsonObject JsonObject) Append() {
	jsonDoc += "\n{"
	for index, element := range jsonObject.Elements {
		if index > 0 {
			jsonDoc += ","
		}
		element.Append()
	}
	jsonDoc += "\n}"
}
