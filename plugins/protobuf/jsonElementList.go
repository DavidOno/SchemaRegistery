package main

type JsonElementList struct {
	JsonElements []Json
}

func (jsonElementList JsonElementList) Append() {
	for index, jsonElement := range jsonElementList.JsonElements {
		if index > 0 {
			jsonDoc += ","
		}
		jsonElement.Append()
	}
}
