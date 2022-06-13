SRC_DIR="./proto"
DST_DIR="./proto"
go install research/grpc/cmd/protoc-gen-go-echo_error && \
protoc            \
-I=$SRC_DIR   $SRC_DIR"/errors.proto"     \
--go_out=$DST_DIR \
--go-echo_error_out=$DST_DIR  \

