dir=$(cd -P -- "$(dirname -- "$0")" && pwd -P)
protoc --proto_path=$dir/proto --micro_out=$dir/proto --go_out=$dir/proto sellinfo.proto