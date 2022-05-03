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
var specifiedFields = map[int]string{
	0: "name",
	1: "optional",
	2: "type",
	3: "typeRef",
	4: "minCardinality",
	5: "maxCardinality"}

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
		messages := flattenMessages(file)
		buf = createJSON(file, messages)

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

func debug(file *protogen.File) {
	fmt.Println("DEBUG: ")
	fmt.Println(file.Messages[0].Desc.Name())
	fmt.Println(file.Messages[0].Messages[0].Desc.Name())
	fmt.Println(file.Messages[0].Enums[0].Desc.Name())
	fmt.Println(file.Messages[1].Desc.Name())
	fmt.Println(file.Enums[0].Desc.Name())
}

func flattenMessages(file *protogen.File) []*protogen.Message {
	var messages []*protogen.Message
	for _, message := range file.Messages {
		messages = append(messages, message)
		for _, nestedMessage := range message.Messages {
			messages = append(messages, nestedMessage)
		}
	}
	return messages
}

func createJSON(file *protogen.File, messages []*protogen.Message) bytes.Buffer {
	// debug(file)
	var buf bytes.Buffer
	root := JsonObject{}
	topLevelList := JsonKVList{}
	schemaName := JsonKV{"name", String{file.GeneratedFilenamePrefix}}
	arrayOfComponents := JsonArray{}
	for _, msg := range messages {
		messageProperties := JsonKVList{}
		messageName := JsonKV{"name", String{string(msg.Desc.Name())}}
		fieldsArray := JsonArray{}
		for _, field := range msg.Fields {
			fieldObj := JsonObject{}
			fieldProperties := JsonKVList{}
			for i := 0; i < len(specifiedFields); i++ {
				specifiedField := JsonKV{specifiedFields[i], String{string(field.Desc.Name())}}
				fieldProperties.JsonElements = append(fieldProperties.JsonElements, specifiedField)
			}
			fieldObj.Elements = append(fieldObj.Elements, fieldProperties)
			fieldsArray.Objects = append(fieldsArray.Objects, fieldObj)
		}
		fields := JsonKV{"fields", fieldsArray}
		messageProperties.JsonElements = append(messageProperties.JsonElements, messageName, fields)
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
