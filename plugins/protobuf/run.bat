go install . 

protoc --proto_path . -I=. test.proto --foo_out=./out --go_out=./out