protoc -I%GOPATH%\src -I%~dp0..\pkg\protobuf\ --go_out=%~dp0..\pkg\protobuf --go-grpc_out=require_unimplemented_servers=false:%~dp0..\pkg\protobuf %~dp0..\pkg\protobuf\*.proto
