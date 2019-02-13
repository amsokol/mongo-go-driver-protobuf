package codecs

import (
	"reflect"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/amsokol/mongo-go-driver-protobuf/mongodb"
)

func TestCodecs(t *testing.T) {
	rb := bson.NewRegistryBuilder()
	r := Register(rb).Build()

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
		Timestamp   *timestamp.Timestamp
		ObjectID    *mongodb.ObjectId
	}

	t.Run("marshal/unmarshal", func(t *testing.T) {
		tm := time.Now()
		// BSON accuracy is in milliseconds
		tm = time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second(),
			(tm.Nanosecond()/1000000)*1000000, tm.Location())

		ts, err := ptypes.TimestampProto(tm)
		if err != nil {
			t.Errorf("ptypes.TimestampProto error = %v", err)
			return
		}

		id := mongodb.ObjectId{Id: "5c601716e1f2d109887d6db2"}

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
			Timestamp:   ts,
			ObjectID:    &id,
		}

		b, err := bson.MarshalWithRegistry(r, &in)
		if err != nil {
			t.Errorf("bson.MarshalWithRegistry error = %v", err)
			return
		}

		var out Data

		if err = bson.UnmarshalWithRegistry(r, b, &out); err != nil {
			t.Errorf("bson.UnmarshalWithRegistry error = %v", err)
			return
		}

		if !reflect.DeepEqual(in, out) {
			t.Errorf("failed: in=%#v, out=%#v", in, out)
		}
	})
}
