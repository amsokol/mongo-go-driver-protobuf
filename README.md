# mongo-go-driver-protobuf

This is extension for officional MongoDB Go driver adds support for Google protocol buffers types.

- [Description](#description)
- [Links](#links)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage example](#usage-example)

## Description

It contains set of BSON marshal/unmarshal codecs for Google protocol buffers type wrappers and Timestamp type:

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

## Links

- Official MongoDB Go Driver: [https://github.com/mongodb/mongo-go-driver](https://github.com/mongodb/mongo-go-driver)
- Google protocol buffers types (wrappers): [https://github.com/golang/protobuf/blob/master/ptypes/wrappers/wrappers.proto](https://github.com/golang/protobuf/blob/master/ptypes/wrappers/wrappers.proto)
- Google protocol buffers Timestamp type: [https://github.com/golang/protobuf/blob/master/ptypes/timestamp/timestamp.proto](https://github.com/golang/protobuf/blob/master/ptypes/timestamp/timestamp.proto)
  
## Requirements

- Google protocol buffers version `proto3`
- Official MongoDB Go Driver Beta 1 or higher

## Installation

Installing using `go get`:

```bash
go get -u github.com/amsokol/mongo-go-driver-protobuf
```

or you don't need to do anything manually if you are using Go modules. Go modules installs necessary packages automatically.

## Usage example

1. Install MongoDB server locally
2. Allow MongoDB server login without login/password (*for test purpose only - don't do it in production!*)
3. Create `experiments` database
4. Create `proto` collection into `experiments` database
5. Take a look and run example:

```go
package main

import (
    "context"
    "log"
    "time"

    "github.com/golang/protobuf/ptypes/timestamp"
    "github.com/golang/protobuf/ptypes"
    "github.com/golang/protobuf/ptypes/wrappers"

    "github.com/mongodb/mongo-go-driver/bson"
    "github.com/mongodb/mongo-go-driver/bson/bsoncodec"
    "github.com/mongodb/mongo-go-driver/bson/primitive"
    "github.com/mongodb/mongo-go-driver/mongo"
    "github.com/mongodb/mongo-go-driver/mongo/options"
    "github.com/mongodb/mongo-go-driver/x/mongo/driver/topology"

    "github.com/amsokol/mongo-go-driver-protobuf"
)

type Data struct {
    BoolValue   *wrappers.BoolValue
    BytesValue  *wrappers.BytesValue
    DoubleValue *wrappers.DoubleValue
    FloatValue  *wrappers.FloatValue
    Int32Value  *wrappers.Int32Value
    Int64Value  *wrappers.Int64Value
    StringValue *wrappers.StringValue
    Uint32Value *wrappers.UInt32Value
    Uint64Value *wrappers.UInt64Value

    Timestamp *timestamp.Timestamp
}

func main() {
    log.Printf("connecting to MongoDB server...")

    rb := bson.NewRegistryBuilder()
    rb = codecs.Register(rb)
    reg := rb.Build()
    client, err := mongo.NewClientWithOptions("mongodb://localhost:27017",
        &options.ClientOptions{
            Registry: reg,
            TopologyOptions: []topology.Option{
                topology.WithServerOptions(func(opts ...topology.ServerOption) []topology.ServerOption {
                    return []topology.ServerOption{
                        topology.WithRegistry(func(r *bsoncodec.Registry) *bsoncodec.Registry {
                            return reg
                        }),
                    }
                }),
            },
        })
    if err != nil {
        log.Fatalf("failed to create new MongoDB client: %#v", err)
    }
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatalf("failed to connect to MongoDB: %#v", err)
    }

    log.Printf("done")

    // get collection
    coll := client.Database("experiments").Collection("proto")

    t := time.Now()
    ts, _ := ptypes.TimestampProto(t)

    in := Data{
        BoolValue:   &wrappers.BoolValue{Value: true},
        BytesValue:  &wrappers.BytesValue{Value: make([]byte, 5)},
        DoubleValue: &wrappers.DoubleValue{Value: 1.2},
        FloatValue:  &wrappers.FloatValue{Value: 1.3},
        Int32Value:  &wrappers.Int32Value{Value: -12345},
        Int64Value:  &wrappers.Int64Value{Value: -123456789},
        StringValue: &wrappers.StringValue{Value: "qwerty"},
        Uint32Value: &wrappers.UInt32Value{Value: 12345},
        Uint64Value: &wrappers.UInt64Value{Value: 123456789},

        Timestamp: ts,
    }

    log.Printf("insert data into collection <experiments.proto>...")
    res, err := coll.InsertOne(ctx, &in)
    if err != nil {
        log.Fatalf("insert data into collection <experiments.proto>: %#v", err)
    }
    id := res.InsertedID
    log.Printf("done, id=%s", id.(primitive.ObjectID).Hex())

    var out Data
    filter := bson.D{{Key: "_id", Value: id}}
    cur := coll.FindOne(ctx, filter)

    log.Printf("selecting data with id=%s from  collection <experiments.proto>...", id.(primitive.ObjectID).Hex())
    err = cur.Decode(&out)
    if err != nil {
        log.Fatalf("failed to get data (id=%#v) from collection <experiments.proto>: %#v", id, err)
    }

    log.Printf("done")
}
```

> *Note*: the following code is mandatory for MongoDB Go Driver Beta 1 only.
> It is removed for new coming driver version:

```go
            TopologyOptions: []topology.Option{
                topology.WithServerOptions(func(opts ...topology.ServerOption) []topology.ServerOption {
                    return []topology.ServerOption{
                        topology.WithRegistry(func(r *bsoncodec.Registry) *bsoncodec.Registry {
                            return reg
                        }),
                    }
                }),
            },
```
