package main

type JsonKVList struct {
	JsonElements []JsonElement
}

func (jsonElementList JsonKVList) Append() {
	for index, jsonElement := range jsonElementList.JsonElements {
		if index > 0 {
			jsonDoc += ","
		}
		jsonElement.Append()
	}
}
