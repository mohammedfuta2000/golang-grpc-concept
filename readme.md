nano ~/.bashrc

add this
export PATH=$PATH:~/go/bin

source ~/.bashrc

go env


/Users/mohammedfuta/go/bin/protoc-gen-go
/Users/mohammedfuta/go/bin/protoc-gen-go-grpc

export PATH=$PATH:/Users/mohammedfuta/go/bin/protoc-gen-go
export PATH=$PATH:/Users/mohammedfuta/go/bin/protoc-gen-go-grpc

/usr/local/Cellar/protobuf@3/3.20.3/bin/protoc --go_out=. --go_grpc_out=. proto/greet.proto