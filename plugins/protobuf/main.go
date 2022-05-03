package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
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

func getType(kind protoreflect.Kind) string {
	// switch kind {
	// case 1:
	// 	return "double"
	// case 2:
	// 	return "float"
	// case 3:
	// 	return "int64"
	// case 4:
	// 	return "uint64"
	// case 5:
	// 	return "int32"
	// case 6:
	// 	return "fixed64"
	// case 7:
	// 	return "fixed32"
	// case 8:
	// 	return "boolean"
	// case 9:
	// 	return "string"
	// case 12:
	// 	return "byte"
	// case 13:
	// 	return "uint32"
	// case 14:
	// 	return "enum"
	// case 15:
	// 	return "fixed32"
	// case 16:
	// 	return "fixed64"
	// case 17:
	// 	return "int32"
	// default:
	// 	return "unknown type - this should not happen"
	switch kind {
	case protoreflect.BoolKind:
		return "boolean"
	case protoreflect.EnumKind:
		return "enum"
	case protoreflect.Int32Kind:
		return "int32"
	case protoreflect.Sint32Kind:
		return "int32"
	case protoreflect.Uint32Kind:
		return "uint32"
	case protoreflect.Int64Kind:
		return "int64"
	case protoreflect.Sint64Kind:
		return "int64"
	case protoreflect.Uint64Kind:
		return "uint64"
	case protoreflect.Sfixed32Kind:
		return "fixed32"
	case protoreflect.Fixed32Kind:
		return "fixed32"
	case protoreflect.FloatKind:
		return "float"
	case protoreflect.Sfixed64Kind:
		return "fixed64"
	case protoreflect.Fixed64Kind:
		return "fixed64"
	case protoreflect.DoubleKind:
		return "double"
	case protoreflect.StringKind:
		return "string"
	case protoreflect.BytesKind:
		return "byte"
	case protoreflect.MessageKind:
		return "not supported"
	case protoreflect.GroupKind:
		return "not supported"
	default:
		return "Error: unknown type of field"
	}

}

func getIfOptional(cardinality protoreflect.Cardinality) string {
	switch cardinality {
	case protoreflect.Optional:
		return "true"
	case protoreflect.Required:
		return "false"
	case protoreflect.Repeated:
		return "false" // appears zero(emptyList) or more times
	default:
		return "Error: unknown if optional, required or repated"
	}
}

func getMinCardinality(cardinality protoreflect.Cardinality) string {
	switch cardinality {
	case protoreflect.Optional:
		return "1"
	case protoreflect.Required:
		return "1"
	case protoreflect.Repeated:
		return "0" // appears zero(emptyList) or more times
	default:
		return "Error: unknown min cardinality"
	}
}

func getMaxCardinality(cardinality protoreflect.Cardinality) string {
	switch cardinality {
	case protoreflect.Optional:
		return "1"
	case protoreflect.Required:
		return "1"
	case protoreflect.Repeated:
		return "*" // appears zero(emptyList) or more times
	default:
		return "Error: unknown max cardinality"
	}
}

func createJSON(file *protogen.File, messages []*protogen.Message) bytes.Buffer {
	// debug(file)
	var buf bytes.Buffer
	root := JsonObject{}
	topLevelList := JsonKVList{}
	schemaName := JsonKV{"name", String{file.GeneratedFilenamePrefix}}
	arrayOfComponents := JsonArray{}
	addMessages(messages, &arrayOfComponents)
	components := JsonKV{"components", arrayOfComponents}
	topLevelList.JsonElements = append(topLevelList.JsonElements, schemaName, components)
	root.Elements = append(root.Elements, topLevelList)
	root.Append(0)
	buf.Write([]byte(jsonDoc))
	return buf
}

func addMessages(messages []*protogen.Message, arrayOfComponents *JsonArray) {
	for _, msg := range messages {
		messageProperties := JsonKVList{}
		messageName := JsonKV{"name", String{string(msg.Desc.Name())}}
		fieldsArray := JsonArray{}
		addFields(msg, &fieldsArray)
		fields := JsonKV{"fields", fieldsArray}
		messageProperties.JsonElements = append(messageProperties.JsonElements, messageName, fields)
		messageObject := JsonObject{}
		messageObject.Elements = append(messageObject.Elements, messageProperties)
		message := JsonKV{"object", messageObject}
		messageWrapperObj := JsonObject{}
		messageWrapperObj.Elements = append(messageWrapperObj.Elements, message)
		arrayOfComponents.Objects = append(arrayOfComponents.Objects, messageWrapperObj)
	}
}

func addFields(msg *protogen.Message, fieldsArray *JsonArray) {
	for _, field := range msg.Fields {
		fieldObj := JsonObject{}
		fieldProperties := JsonKVList{}
		for i := 0; i < len(specifiedFields); i++ {
			specifiedField := JsonKV{}
			specifiedField.Name = specifiedFields[i]
			switch i {
			case 0:
				specifiedField.Value = String{string(field.Desc.Name())}
			case 1:
				specifiedField.Value = String{getIfOptional(field.Desc.Cardinality())}
			case 2:
				specifiedField.Value = String{getType(field.Desc.Kind())}
			case 4:
				specifiedField.Value = String{getMinCardinality(field.Desc.Cardinality())}
			case 5:
				specifiedField.Value = String{getMaxCardinality(field.Desc.Cardinality())}
			default:
				specifiedField.Value = String{string(field.Desc.Name())}
			}
			fieldProperties.JsonElements = append(fieldProperties.JsonElements, specifiedField)
		}
		fieldObj.Elements = append(fieldObj.Elements, fieldProperties)
		fieldsArray.Objects = append(fieldsArray.Objects, fieldObj)
	}
}
