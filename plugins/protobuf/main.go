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

var jsonDoc string = ""

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

func createJSON(file *protogen.File) bytes.Buffer {

}

func createJson(file *protogen.File) bytes.Buffer {
	var buf bytes.Buffer
	addJSONBeginning()
	addSchemaName(file)
	addComponents(file)
	addJSONEnding()
	buf.Write([]byte(jsonDoc))
	return buf
}

func addJSONBeginning() {
	jsonDoc += "{"
}

func addSchemaName(file *protogen.File) {
	newLine("name", file.GeneratedFilenamePrefix, 0)
}

func addComponents(file *protogen.File) {
	newLineOnlyTag("components", "[{", 0)
	for index, msg := range file.Proto.MessageType {
		if index > 0 {
			jsonDoc += ","
		}
		newLineOnlyTag("object", "{", 1)
		newLine("name", *msg.Name, 2)
		newLineSingleElement("}", 1)
	}
	newLineSingleElement("]", 0)
}

func addJSONEnding() {
	newLineSingleElement("}", 0)
}

func newLine(toAddTag string, toAddValue string, level int) {
	jsonDoc += "\n" + addTabs(level) + fmt.Sprintf("\"%s\": \"%s\",", toAddTag, toAddValue)
}

func newLineOnlyTag(tagToAdd string, remainderToAdd string, level int) {
	jsonDoc += "\n" + addTabs(level) + fmt.Sprintf("\"%s\": %s", tagToAdd, remainderToAdd)
}

func newLineSingleElement(toAdd string, level int) {
	jsonDoc += "\n" + addTabs(level) + toAdd
}

func addTabs(level int) string {
	tabs := ""
	for i := 0; i < level; i++ {
		tabs += "\t"
	}
	return tabs
}
