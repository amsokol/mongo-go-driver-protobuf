@protoc --proto_path=proto/mongodb --go_out=../../../ objectid.proto

@protoc --proto_path=./ --proto_path=proto --proto_path=proto/third_party --go_out=. ./codecs_test.proto
