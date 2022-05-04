package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/pluginpb"
)

var jsonDoc string = ""
var specifiedFieldsForMessages = map[int]string{
	0: "name",
	1: "optional",
	2: "type",
	3: "typeRef",
	4: "minCardinality",
	5: "maxCardinality",
	6: "comment"}

var specifiedFieldsForEnums = map[int]string{
	0: "name",
	1: "comment"}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	var req pluginpb.CodeGeneratorRequest
	proto.Unmarshal(input, &req)
	opts := protogen.Options{}
	plugin, err := opts.New(&req)
	if err != nil {
		panic(err)
	}
	for _, file := range plugin.Files {
		var buf bytes.Buffer
		messages := collectMessages(file)
		enums := collectEnums(file)
		buf = createJSON(file, messages, enums)
		filename := file.GeneratedFilenamePrefix + ".json"
		file := plugin.NewGeneratedFile(filename, ".")
		file.Write(buf.Bytes())
	}

	stdout := plugin.Response()
	out, err := proto.Marshal(stdout)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, string(out))
}

func collectMessages(file *protogen.File) []*protogen.Message {
	var messages []*protogen.Message
	for _, message := range file.Messages {
		messages = append(messages, message)
		for _, nestedMessage := range message.Messages {
			messages = append(messages, nestedMessage)
		}
	}
	return messages
}

func collectEnums(file *protogen.File) []*protogen.Enum {
	var enums []*protogen.Enum
	for _, enum := range file.Enums {
		enums = append(enums, enum)
	}
	for _, message := range file.Messages {
		for _, enum := range message.Enums {
			enums = append(enums, enum)
		}
	}
	return enums
}

func getType(kind protoreflect.Kind) string {
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
		return "object"
	case protoreflect.GroupKind:
		return "not supported"
	default:
		return "Error: unknown type of field"
	}
}

func getTypeRef(field protogen.Field) JsonElement {
	switch field.Desc.Kind() {
	case protoreflect.MessageKind:
		return String{resolveReference(field)}
	default:
		return Null{}
	}
}

func resolveReference(field protogen.Field) string {
	return string(field.Desc.Message().FullName())
}

func getIfOptional(cardinality protoreflect.Cardinality) Boolean {
	switch cardinality {
	case protoreflect.Optional:
		return Boolean{true}
	case protoreflect.Required:
		return Boolean{false}
	case protoreflect.Repeated:
		return Boolean{false} // appears zero(emptyList) or more times
	default:
		panic("Error: unknown if optional, required or repated")
	}
}

func getMinCardinality(cardinality protoreflect.Cardinality) JsonElement {
	switch cardinality {
	case protoreflect.Optional:
		return Null{}
	case protoreflect.Required:
		return Null{}
	case protoreflect.Repeated:
		return String{"0"} // appears zero(emptyList) or more times
	default:
		panic("Error: unknown min cardinality")
	}
}

func getMaxCardinality(cardinality protoreflect.Cardinality) JsonElement {
	switch cardinality {
	case protoreflect.Optional:
		return Null{}
	case protoreflect.Required:
		return Null{}
	case protoreflect.Repeated:
		return String{"*"} // appears zero(emptyList) or more times
	default:
		panic("Error: unknown max cardinality")
	}
}

func createJSON(file *protogen.File, messages []*protogen.Message, enums []*protogen.Enum) bytes.Buffer {
	var buf bytes.Buffer
	root := JsonObject{}
	topLevelList := JsonKVList{}
	schemaName := JsonKV{"name", String{file.GeneratedFilenamePrefix}}
	schemaSpecification := JsonKV{"schemaSpec", String{file.Desc.Syntax().String()}}
	comment := JsonKV{"comment", String{findTopLevelComment(file)}}
	arrayOfComponents := JsonArray{}
	addMessages(messages, &arrayOfComponents)
	addEnums(enums, &arrayOfComponents)
	components := JsonKV{"components", arrayOfComponents}
	topLevelList.JsonElements = append(topLevelList.JsonElements, schemaName, schemaSpecification, comment, components)
	root.Elements = append(root.Elements, topLevelList)
	root.Append(0)
	buf.Write([]byte(jsonDoc))
	return buf
}

func addMessages(messages []*protogen.Message, arrayOfComponents *JsonArray) {
	for _, msg := range messages {
		messageProperties := JsonKVList{}
		messageName := JsonKV{"name", String{string(msg.Desc.Name())}}
		comment := JsonKV{"comment", String{extractCommentForMessage(msg)}}
		fieldsArray := JsonArray{}
		addFieldsForMessage(msg, &fieldsArray)
		fields := JsonKV{"fields", fieldsArray}
		messageProperties.JsonElements = append(messageProperties.JsonElements, messageName, comment, fields)
		messageObject := JsonObject{}
		messageObject.Elements = append(messageObject.Elements, messageProperties)
		message := JsonKV{"object", messageObject}
		messageWrapperObj := JsonObject{}
		messageWrapperObj.Elements = append(messageWrapperObj.Elements, message)
		arrayOfComponents.Objects = append(arrayOfComponents.Objects, messageWrapperObj)
	}
}

func addEnums(enums []*protogen.Enum, arrayOfComponents *JsonArray) {
	for _, enum := range enums {
		enumProperties := JsonKVList{}
		enumName := JsonKV{"name", String{string(enum.Desc.Name())}}
		comment := JsonKV{"comment", String{extractCommentForEnum(enum)}}
		fieldsArray := JsonArray{}
		addFieldsForEnum(enum, &fieldsArray)
		fields := JsonKV{"values", fieldsArray}
		enumProperties.JsonElements = append(enumProperties.JsonElements, enumName, comment, fields)
		enumObject := JsonObject{}
		enumObject.Elements = append(enumObject.Elements, enumProperties)
		enum := JsonKV{"enum", enumObject}
		enumWrapperObj := JsonObject{}
		enumWrapperObj.Elements = append(enumWrapperObj.Elements, enum)
		arrayOfComponents.Objects = append(arrayOfComponents.Objects, enumWrapperObj)
	}
}

func addFieldsForMessage(msg *protogen.Message, fieldsArray *JsonArray) {
	for _, field := range msg.Fields {
		fieldObj := JsonObject{}
		fieldProperties := JsonKVList{}
		addFieldProperties(field, &fieldProperties)
		fieldObj.Elements = append(fieldObj.Elements, fieldProperties)
		fieldsArray.Objects = append(fieldsArray.Objects, fieldObj)
	}
}

func addFieldProperties(field *protogen.Field, fieldProperties *JsonKVList) {
	for i := 0; i < len(specifiedFieldsForMessages); i++ {
		specifiedField := JsonKV{}
		specifiedField.Name = specifiedFieldsForMessages[i]
		switch i {
		case 0:
			specifiedField.Value = String{string(field.Desc.Name())}
		case 1:
			specifiedField.Value = getIfOptional(field.Desc.Cardinality())
		case 2:
			specifiedField.Value = String{getType(field.Desc.Kind())}
		case 3:
			specifiedField.Value = getTypeRef(*field)
		case 4:
			specifiedField.Value = getMinCardinality(field.Desc.Cardinality())
		case 5:
			specifiedField.Value = getMaxCardinality(field.Desc.Cardinality())
		case 6:
			specifiedField.Value = String{extractCommentForField(field)}
		default:
			specifiedField.Value = String{string(field.Desc.Name())}
		}
		fieldProperties.JsonElements = append(fieldProperties.JsonElements, specifiedField)
	}
}

func addFieldsForEnum(enum *protogen.Enum, fieldsArray *JsonArray) {
	for _, value := range enum.Values {
		valueList := JsonKVList{}
		valueObj := JsonObject{}
		for i := 0; i < len(specifiedFieldsForEnums); i++ {
			keyValuePair := JsonKV{}
			keyValuePair.Name = specifiedFieldsForEnums[i]
			switch i {
			case 0:
				keyValuePair.Value = String{string(value.Desc.Name())}
			case 1:
				keyValuePair.Value = String{prepareComment(value.Comments.Leading.String())}
			}
			valueList.JsonElements = append(valueList.JsonElements, keyValuePair)
		}
		valueObj.Elements = append(valueObj.Elements, valueList)
		fieldsArray.Objects = append(fieldsArray.Objects, valueObj)
	}

}

func extractCommentForMessage(msg *protogen.Message) string {
	comment := msg.Comments.Leading.String()
	return prepareComment(comment)
}

func extractCommentForEnum(enum *protogen.Enum) string {
	comment := enum.Comments.Leading.String()
	return prepareComment(comment)
}

func extractCommentForField(field *protogen.Field) string {
	comment := field.Comments.Leading.String()
	return prepareComment(comment)
}

func prepareComment(comment string) string {
	comment = removeAllDoubleSlashes(comment)
	comment = removeLastNewLine(comment)
	comment = replaceIntermediateLineBreaks(comment)
	comment = removeAllDoubleWhiteSpaces(comment)
	return strings.Trim(comment, " ")
}

func replaceIntermediateLineBreaks(comment string) string {
	return strings.ReplaceAll(comment, "\n", " ")
}

func removeLastNewLine(comment string) string {
	if strings.HasSuffix(comment, "\n") {
		return comment[:len(comment)-1]
	}
	return comment
}

func removeAllDoubleSlashes(comment string) string {
	return strings.ReplaceAll(comment, "//", "")
}

func removeAllDoubleWhiteSpaces(comment string) string {
	return strings.ReplaceAll(comment, "  ", " ")
}

func findTopLevelComment(file *protogen.File) string {
	for _, location := range file.Proto.SourceCodeInfo.Location {
		if len(location.LeadingDetachedComments) > 0 {
			return prepareComment(location.LeadingDetachedComments[0])
		}
	}
	return ""
}
