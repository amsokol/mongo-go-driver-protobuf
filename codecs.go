package codecs

import (
	"reflect"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/mongodb/mongo-go-driver/bson/bsoncodec"
	"github.com/mongodb/mongo-go-driver/bson/bsonrw"
)

var (
	// Protobuf wrappers types
	boolValueType   = reflect.TypeOf(wrappers.BoolValue{})
	bytesValueType  = reflect.TypeOf(wrappers.BytesValue{})
	doubleValueType = reflect.TypeOf(wrappers.DoubleValue{})
	floatValueType  = reflect.TypeOf(wrappers.FloatValue{})
	int32ValueType  = reflect.TypeOf(wrappers.Int32Value{})
	int64ValueType  = reflect.TypeOf(wrappers.Int64Value{})
	stringValueType = reflect.TypeOf(wrappers.StringValue{})
	uint32ValueType = reflect.TypeOf(wrappers.UInt32Value{})
	uint64ValueType = reflect.TypeOf(wrappers.UInt64Value{})

	// Protobuf Timestamp type
	timestampType = reflect.TypeOf(timestamp.Timestamp{})

	// Time type
	timeType = reflect.TypeOf(time.Time{})

	// Codecs
	wrapperValueCodec = &WrapperValueCodec{}
	timestampCodec    = &TimestampCodec{}
)

// WrapperValueCodec is codec for Protobuf type wrappers
type WrapperValueCodec struct {
}

// EncodeValue encodes Protobuf type wrapper value to BSON value
func (e *WrapperValueCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	val = val.FieldByName("Value")
	enc, err := ectx.LookupEncoder(val.Type())
	if err != nil {
		return err
	}
	return enc.EncodeValue(ectx, vw, val)
}

// DecodeValue decodes BSON value to Protobuf type wrapper value
func (e *WrapperValueCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	val = val.FieldByName("Value")
	enc, err := ectx.LookupDecoder(val.Type())
	if err != nil {
		return err
	}
	return enc.DecodeValue(ectx, vr, val)
}

// TimestampCodec is codec for Protobuf Timestamp
type TimestampCodec struct {
}

// EncodeValue encodes Protobuf Timestamp value to BSON value
func (e *TimestampCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	v := val.Interface().(timestamp.Timestamp)
	t, err := ptypes.Timestamp(&v)
	if err != nil {
		return err
	}
	enc, err := ectx.LookupEncoder(timeType)
	if err != nil {
		return err
	}
	return enc.EncodeValue(ectx, vw, reflect.ValueOf(t.In(time.UTC)))
}

// DecodeValue decodes BSON value to Timestamp value
func (e *TimestampCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	enc, err := ectx.LookupDecoder(timeType)
	if err != nil {
		return err
	}
	var t time.Time
	if err = enc.DecodeValue(ectx, vr, reflect.ValueOf(&t).Elem()); err != nil {
		return err
	}
	ts, err := ptypes.TimestampProto(t.In(time.UTC))
	if err != nil {
		return err
	}
	val.Set(reflect.ValueOf(*ts))
	return nil
}

// Register registers Google protocol buffers types codecs
func Register(rb *bsoncodec.RegistryBuilder) *bsoncodec.RegistryBuilder {
	return rb.RegisterCodec(boolValueType, wrapperValueCodec).
		RegisterCodec(bytesValueType, wrapperValueCodec).
		RegisterCodec(doubleValueType, wrapperValueCodec).
		RegisterCodec(floatValueType, wrapperValueCodec).
		RegisterCodec(int32ValueType, wrapperValueCodec).
		RegisterCodec(int64ValueType, wrapperValueCodec).
		RegisterCodec(stringValueType, wrapperValueCodec).
		RegisterCodec(uint32ValueType, wrapperValueCodec).
		RegisterCodec(uint64ValueType, wrapperValueCodec).
		RegisterCodec(timestampType, timestampCodec)
}
