# mongo-go-driver-protobuf

This is extension for officional MongoDB Go driver adds support for Google protocol buffers types.

- [Description](#description)
- [Links](#links)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage example](#usage-example)

## Description

It contains set of BSON marshal/unmarshal codecs for Google protocol buffers type wrappers, Timestamp and MongoDB ObjectID:

- `BoolValue`
- `BytesValue`
- `DoubleValue`
- `FloatValue`
- `Int32Value`
- `Int64Value`
- `StringValue`
- `Uint32Value`
- `Uint64Value`
- `Timestamp`
- `ObjectID`

## Links

- Official MongoDB Go Driver: [https://github.com/mongodb/mongo-go-driver](https://github.com/mongodb/mongo-go-driver)
- Google protocol buffers types (wrappers): [https://github.com/golang/protobuf/blob/master/ptypes/wrappers/wrappers.proto](https://github.com/golang/protobuf/blob/master/ptypes/wrappers/wrappers.proto)
- Google protocol buffers Timestamp type: [https://github.com/golang/protobuf/blob/master/ptypes/timestamp/timestamp.proto](https://github.com/golang/protobuf/blob/master/ptypes/timestamp/timestamp.proto)
- MongoDB ObjectID type: [https://github.com/mongodb/mongo-go-driver/blob/master/bson/primitive/objectid.go](https://github.com/mongodb/mongo-go-driver/blob/master/bson/primitive/objectid.go)
- MongoDB ObjectID my proto wrapper: [https://github.com/amsokol/mongo-go-driver-protobuf/blob/master/proto/mongodb/objectid.proto](https://github.com/amsokol/mongo-go-driver-protobuf/blob/master/proto/mongodb/objectid.proto)
  
## Requirements

- Google protocol buffers version `proto3`
- Official MongoDB Go Driver Beta 2 or higher

## Installation

Installing using `go get`:

```bash
go get -u github.com/amsokol/mongo-go-driver-protobuf
```

or you don't need to do anything manually if you are using Go modules. Go modules installs necessary packages automatically.

## Usage example

First install `protoc-gen-gotag` to make available Go language `tags` for proto messages

```bash
go get -u github.com/amsokol/protoc-gen-gotag
```

Next

1. Create free Altas mini MongoDB instance
2. Create `experiments` database
3. Create `proto` collection into `experiments` database
4. Run this [example](https://github.com/amsokol/mongo-go-driver-protobuf/tree/master/examples/mongodb-rw)
