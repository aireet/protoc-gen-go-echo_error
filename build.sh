SRC_DIR="./proto"
DST_DIR="./proto"
go install github.com/aireet/protoc-gen-go-echo_error && \
protoc            \
-I=$SRC_DIR   $SRC_DIR"/errors.proto"     \
--go_out=$DST_DIR \
--go-echo_error_out=$DST_DIR  \

