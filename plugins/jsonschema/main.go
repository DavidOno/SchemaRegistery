package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	input, err := os.ReadFile("./test.json")
	check(err)

	// Declared an empty map interface
	var result map[string]interface{}

	// Unmarshal or Decode the JSON to the interface.
	json.Unmarshal(input, &result)

	//Transform

	//Write in file
	f, err := os.Create("jss2dc_test.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	bytes, _ := json.MarshalIndent(result, "", "\t")
	_, err2 := f.Write(bytes)
	if err2 != nil {
		log.Fatal(err2)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
