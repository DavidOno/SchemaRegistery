go install . 

protoc --proto_path . -I=. test.proto --datacatalog_out=./out --go_out=./out