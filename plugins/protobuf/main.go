package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	// Protoc passes pluginpb.CodeGeneratorRequest in via stdin
	// marshalled with Protobuf
	input, _ := ioutil.ReadAll(os.Stdin)
	var req pluginpb.CodeGeneratorRequest
	proto.Unmarshal(input, &req)

	// Initialise our plugin with default options
	opts := protogen.Options{}
	plugin, err := opts.New(&req)
	if err != nil {
		panic(err)
	}

	// Protoc passes a slice of File structs for us to process
	for _, file := range plugin.Files {

		// Time to generate code...!

		// 1. Initialise a buffer to hold the generated code
		var buf bytes.Buffer
		buf = createJson(file)

		// 4. Specify the output filename, in this case test.foo.go
		filename := file.GeneratedFilenamePrefix + ".json"
		file := plugin.NewGeneratedFile(filename, ".")

		// 5. Pass the data from our buffer to the plugin file struct
		file.Write(buf.Bytes())
	}

	// Generate a response from our plugin and marshall as protobuf
	stdout := plugin.Response()
	out, err := proto.Marshal(stdout)
	if err != nil {
		panic(err)
	}

	// Write the response to stdout, to be picked up by protoc
	fmt.Fprintf(os.Stdout, string(out))
}

func createJson(file *protogen.File) bytes.Buffer {
	var buf bytes.Buffer
	jsonDoc := ""
	addJSONBeginning(&jsonDoc)
	addSchemaName(&jsonDoc, file)
	addComponents(&jsonDoc, file)
	addJSONEnding(&jsonDoc)
	buf.Write([]byte(jsonDoc))
	return buf
}

func addJSONBeginning(jsonDoc *string) {
	*jsonDoc += "{"
}

func addSchemaName(jsonDoc *string, file *protogen.File) {
	newLine(jsonDoc, "name", file.GeneratedFilenamePrefix, 0)
}

func addComponents(jsonDoc *string, file *protogen.File) string {
	components := ""
	newLineOnlyTag(jsonDoc, "components", "[{", 0)
	for index, msg := range file.Proto.MessageType {
		component := ""
		if index > 0 {
			component += ","
		}
		newLineOnlyTag(jsonDoc, "object", "{", 1)
		newLine(jsonDoc, "name", *msg.Name, 2)
		newLineSingleElement(jsonDoc, "}", 1)
	}
	newLineSingleElement(jsonDoc, "]", 0)
	return components
}

func addJSONEnding(jsonDoc *string) {
	newLineSingleElement(jsonDoc, "}", 0)
}

func newLine(toAddTo *string, toAddTag string, toAddValue string, level int) {
	*toAddTo += "\n" + addTabs(level) + fmt.Sprintf("\"%s\": \"%s\",", toAddTag, toAddValue)
}

func newLineOnlyTag(toAddTo *string, tagToAdd string, remainderToAdd string, level int) {
	*toAddTo += "\n" + addTabs(level) + fmt.Sprintf("\"%s\": %s", tagToAdd, remainderToAdd)
}

func newLineSingleElement(toAddTo *string, toAdd string, level int) {
	*toAddTo += "\n" + addTabs(level) + toAdd
}

func addTabs(level int) string {
	tabs := ""
	for i := 0; i < level; i++ {
		tabs += "\t"
	}
	return tabs
}
