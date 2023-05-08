pb:
	if [ -d "pb" ]; then rm -rf pb; fi && mkdir pb

server_proto:
	if [ -d "server/pb" ]; then rm -rf server/pb; fi && mkdir server/pb && protoc --go_out=./server/pb --go-grpc_out=./server/pb proto/*.proto 

client_proto:
	if [ -d "client/pb" ]; then rm -rf client/pb; fi && mkdir client/pb && protoc --go_out=./client/pb --go-grpc_out=./client/pb proto/*.proto 

proto:
	if [ -d "gen/api/v1" ]; then rm -rf gen/api/v1; fi && protoc --go_out=./gen --go-grpc_out=./gen api/v1/*.proto 

test:
	protoc -Iproto/ --go_out=. --go_opt=module=github.com/azar-intelops/grpc-code-gen --go-grpc_out=. --go-grpc_opt=module=github.com/azar-intelops/grpc-code-gen proto/*.proto

path:
	export PATH="$PATH:$(go env GOPATH)/bin"