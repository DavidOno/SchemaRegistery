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

		//Intermediate
		// info := fmt.Sprintf("info: \n %v", file)
		// buf.Write([]byte(info))

		// 2. Write the package name
		// 3. For each message add our Foo() method
		// for _, msg := range file.Proto.MessageType {
		//Intermediate - start
		// info := fmt.Sprintf("info: \n %v\n\n", *msg)
		// buf.Write([]byte(info))
		// info = fmt.Sprintf("info2: \n %v \n\n", msg.GetField()[0].Type)
		// buf.Write([]byte(info))
		//Intermediate - end
		// buf.Write([]byte(fmt.Sprintf(`
		// func (x %s) Foo() string {
		//    return "bar"
		// }`, *msg.Name)))
		// }
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
	return fmt.Sprintf("{\n")
}

func addSchemaName(file *protogen.File) string {
	return fmt.Sprintf("\"name\": \"%s\",\n", file.GeneratedFilenamePrefix)
}

func addComponents(file *protogen.File) string {
	components := "\"components\": [\n"
	for index, msg := range file.Proto.MessageType {
		var component string
		if index > 0 {
			component += ","
		}
		component += "{\"object\": {\n"
		component += fmt.Sprintf("\"name\": \"%s\"", *msg.Name)
		component += "}\n}"
		components += component
	}
	components += "]"
	return components
}

func addJSONEnding() string {
	return fmt.Sprintf("\n}")
}
