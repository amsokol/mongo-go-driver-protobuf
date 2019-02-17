@protoc --proto_path=. --proto_path=../../proto --proto_path=../../proto/third_party --proto_path=../third_party --go_out=. data.proto

@protoc --proto_path=. --proto_path=../../proto --proto_path=../../proto/third_party --proto_path=../third_party --gotag_out=xxx="bson+\"-\"",output_path=.:. data.proto
