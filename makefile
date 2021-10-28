.PHONY: install
install:
	go install \
           google.golang.org/protobuf/cmd/protoc-gen-go \
           google.golang.org/grpc/cmd/protoc-gen-go-grpc

.PHONY: create_proto
create_proto: install
	protoc \
    		--go_out=. \
    		--go_opt=module=chatRooms \
    		--go-grpc_out=. \
    		--go-grpc_opt=module=chatRooms \
          	-Iproto \
          			$(shell find proto -type f -name '*.proto')