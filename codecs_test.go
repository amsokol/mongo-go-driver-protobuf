package codecs

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/amsokol/mongo-go-driver-protobuf/mongodb"
	"github.com/amsokol/mongo-go-driver-protobuf/test"
)

func TestCodecs(t *testing.T) {
	rb := bson.NewRegistryBuilder()
	r := Register(rb).Build()

	tm := time.Now()
	// BSON accuracy is in milliseconds
	tm = time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second(),
		(tm.Nanosecond()/1000000)*1000000, tm.Location())

	ts, err := ptypes.TimestampProto(tm)
	if err != nil {
		t.Errorf("ptypes.TimestampProto error = %v", err)
		return
	}

	objectID := primitive.NewObjectID()
	id := mongodb.ObjectId{Value: objectID.Hex()}

	t.Run("primitive object id", func(t *testing.T) {
		resultID, err := id.GetObjectID()
		if err != nil {
			t.Errorf("mongodb.ObjectId.GetPrimitiveObjectID() error = %v", err)
			return
		}

		if !reflect.DeepEqual(objectID, resultID) {
			t.Errorf("failed: primitive object ID=%#v, ID=%#v", objectID, id)
			return
		}
	})

	in := test.Data{
		BoolValue:   &wrappers.BoolValue{Value: true},
		BytesValue:  &wrappers.BytesValue{Value: make([]byte, 5)},
		DoubleValue: &wrappers.DoubleValue{Value: 1.2},
		FloatValue:  &wrappers.FloatValue{Value: 1.3},
		Int32Value:  &wrappers.Int32Value{Value: -12345},
		Int64Value:  &wrappers.Int64Value{Value: -123456789},
		StringValue: &wrappers.StringValue{Value: "qwerty"},
		Uint32Value: &wrappers.UInt32Value{Value: 12345},
		Uint64Value: &wrappers.UInt64Value{Value: 123456789},
		Timestamp:   ts,
		Id:          &id,
	}

	t.Run("marshal/unmarshal", func(t *testing.T) {
		b, err := bson.MarshalWithRegistry(r, &in)
		if err != nil {
			t.Errorf("bson.MarshalWithRegistry error = %v", err)
			return
		}

		var out test.Data

		if err = bson.UnmarshalWithRegistry(r, b, &out); err != nil {
			t.Errorf("bson.UnmarshalWithRegistry error = %v", err)
			return
		}

		if !reflect.DeepEqual(in, out) {
			t.Errorf("failed: in=%#v, out=%#v", in, out)
			return
		}
	})

	t.Run("marshal-jsonpb/unmarshal-jsonpb", func(t *testing.T) {
		var b bytes.Buffer

		m := &jsonpb.Marshaler{}

		if err := m.Marshal(&b, &in); err != nil {
			t.Errorf("jsonpb.Marshaler.Marshal error = %v", err)
			return
		}

		var out test.Data
		if err = jsonpb.Unmarshal(&b, &out); err != nil {
			t.Errorf("jsonpb.Unmarshal error = %v", err)
			return
		}

		if !reflect.DeepEqual(in, out) {
			t.Errorf("failed: in=%#v, out=%#v", in, out)
			return
		}
	})
}
