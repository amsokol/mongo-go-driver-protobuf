@protoc --proto_path=proto/mongodb --go_out=../../../ objectid.proto

@protoc --proto_path=test --proto_path=proto --proto_path=proto/third_party --go_out=test codecs_test.proto
