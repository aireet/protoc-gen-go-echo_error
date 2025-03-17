SRC_DIR="./source"
DST_DIR="./proto"
go22 install . && \

protoc            \
-I $SRC_DIR       \
--go_out=$DST_DIR \
--go_opt paths=source_relative \
--go-echo_error_out=$DST_DIR  \
--go-echo_error_opt require_unimplemented_servers=false,paths=source_relative \
$SRC_DIR/*.proto