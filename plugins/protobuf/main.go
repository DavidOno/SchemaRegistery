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
		buf = createJSON(file)

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
	var buf bytes.Buffer
	root := JsonObject{}
	topLevelList := JsonKVList{}
	schemaName := JsonKV{"name", String{file.GeneratedFilenamePrefix}}
	arrayOfComponents := JsonArray{}
	for _, msg := range file.Proto.MessageType {
		messageProperties := JsonKVList{}
		messageName := JsonKV{"name", String{*msg.Name}}
		messageProperties.JsonElements = append(messageProperties.JsonElements, messageName)
		messageObject := JsonObject{}
		messageObject.Elements = append(messageObject.Elements, messageProperties)
		message := JsonKV{"object", messageObject}
		messageWrapperObj := JsonObject{}
		messageWrapperObj.Elements = append(messageWrapperObj.Elements, message)
		arrayOfComponents.Objects = append(arrayOfComponents.Objects, messageWrapperObj)
	}
	components := JsonKV{"components", arrayOfComponents}
	topLevelList.JsonElements = append(topLevelList.JsonElements, schemaName, components)
	root.Elements = append(root.Elements, topLevelList)
	root.Append(0)
	buf.Write([]byte(jsonDoc))
	return buf
}
