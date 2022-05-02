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
	buf.Write([]byte(addJSONBeginning()))
	buf.Write([]byte(addSchemaName(file)))
	buf.Write([]byte(addComponents(file)))
	buf.Write([]byte(addJSONEnding()))
	return buf
}

func addJSONBeginning() string {
	jsonBegin := ""
	newLineSingleElement(&jsonBegin, "{", 0)
	return jsonBegin
	// return "{\n"
}

func addSchemaName(file *protogen.File) string {
	schemaName := ""
	newLine(&schemaName, "name", file.GeneratedFilenamePrefix, 0)
	return schemaName
	// return fmt.Sprintf("\"name\": \"%s\",\n", file.GeneratedFilenamePrefix)
}

func addComponents(file *protogen.File) string {
	components := ""
	newLineOnlyTag(&components, "components", "[{", 0)
	for index, msg := range file.Proto.MessageType {
		component := ""
		if index > 0 {
			component += ","
		}
		newLineOnlyTag(&component, "object", "{", 1)
		newLine(&component, "name", *msg.Name, 2)
		newLineSingleElement(&component, "}", 1)
		components += component
	}
	components += "]"
	return components
}

func addJSONEnding() string {
	jsonEnd := ""
	newLineSingleElement(&jsonEnd, "}", 0)
	return jsonEnd
	// return fmt.Sprintf("\n}")
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
