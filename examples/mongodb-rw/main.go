package main

import (
	"bytes"
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"

	"github.com/amsokol/mongo-go-driver-protobuf"
)

func main() {
	log.Printf("connecting to MongoDB...")

	// Register custom codecs for protobuf Timestamp and wrapper types
	reg := codecs.Register(bson.NewRegistryBuilder()).Build()

	// Create MongoDB client with registered custom codecs for protobuf Timestamp and wrapper types
	// NOTE: "mongodb+srv" protocol means connect to Altas cloud MongoDB server
	//       use just "mongodb" if you connect to on-premise MongoDB server
	client, err := mongo.NewClientWithOptions("mongodb+srv://USER:PASSWORD@SERVER/experiments",
		&options.ClientOptions{
			Registry: reg,
		})

	if err != nil {
		log.Fatalf("failed to create new MongoDB client: %#v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect client
	if err = client.Connect(ctx); err != nil {
		log.Fatalf("failed to connect to MongoDB: %#v", err)
	}

	log.Printf("connected successfully")

	// Get collection from database
	coll := client.Database("experiments").Collection("proto")

	// Create protobuf Timestamp value from golang Time
	ts := ptypes.TimestampNow()

	// Fill in data structure
	in := Data{
		BoolValue:      true,
		BoolProtoValue: &wrappers.BoolValue{Value: true},

		BytesValue:      []byte{1, 2, 3, 4, 5},
		BytesProtoValue: &wrappers.BytesValue{Value: []byte{1, 2, 3, 4, 5}},

		DoubleValue:      123.45678,
		DoubleProtoValue: &wrappers.DoubleValue{Value: 123.45678},

		FloatValue:      123.45,
		FloatProtoValue: &wrappers.FloatValue{Value: 123.45},

		Int32Value:      -12345,
		Int32ProtoValue: &wrappers.Int32Value{Value: -12345},

		Int64Value:      -123456789000,
		Int64ProtoValue: &wrappers.Int64Value{Value: -123456789000},

		StringValue:      "qwerty",
		StringProtoValue: &wrappers.StringValue{Value: "qwerty"},

		Uint32Value:      12345,
		Uint32ProtoValue: &wrappers.UInt32Value{Value: 12345},

		Uint64Value:      123456789000,
		Uint64ProtoValue: &wrappers.UInt64Value{Value: 123456789000},

		Timestamp: ts,
	}

	log.Printf("insert data into collection <experiments.proto>...")

	// Insert data into the collection
	res, err := coll.InsertOne(ctx, &in)
	if err != nil {
		log.Fatalf("insert data into collection <experiments.proto>: %#v", err)
	}
	id := res.InsertedID
	log.Printf("inserted new item with id=%v successfully", id)

	// Create filter and output structure to read data from collection
	var out Data
	filter := bson.D{{Key: "_id", Value: id}}

	// Read data from collection
	err = coll.FindOne(ctx, filter).Decode(&out)
	if err != nil {
		log.Fatalf("failed to read data (id=%v) from collection <experiments.proto>: %#v", id, err)
	}

	var b bytes.Buffer
	m := &jsonpb.Marshaler{Indent: "  "}
	if err := m.Marshal(&b, &out); err != nil {
		log.Fatalf("jsonpb.Marshaler.Marshal error = %v", err)
	}

	log.Printf("read successfully:\n%s", b.String())
}
