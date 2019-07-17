for file in `find ./srv/ -name "*.proto"`;do
    dir=$(dirname "${file}")
    protoc --proto_path=$dir --micro_out=$dir --go_out=$dir `basename "${file}"`
done
