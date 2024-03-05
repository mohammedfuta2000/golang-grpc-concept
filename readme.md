GRPC rpoject built with googles proto and GRPC package

cuz i messed up my paths
/usr/local/Cellar/protobuf@3/3.20.3/bin/protoc --go_out=. --go_grpc_out=. proto/greet.proto

but this should be fine
protoc --go_out=. --go_grpc_out=. proto/greet.proto

1. write up a proto and define the rpcs u want. 
    - each rpc could be of one of four types: unary, server streaming, client streaming, bi-directional
    - then generate it
2. then build server
3. build cleint
    - for corresponding rpc type(unary, server_streaming, cleint_streaming, bi_sirectional_streaming)