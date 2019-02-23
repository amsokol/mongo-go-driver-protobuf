package codecs

import (
	"reflect"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/amsokol/mongo-go-driver-protobuf/mongodb"
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

	// ObjectId type
	objectIDType          = reflect.TypeOf(mongodb.ObjectId{})
	objectIDPrimitiveType = reflect.TypeOf(primitive.ObjectID{})

	// Codecs
	wrapperValueCodecRef = &wrapperValueCodec{}
	timestampCodecRef    = &timestampCodec{}
	objectIDCodecRef     = &objectIDCodec{}
)

// wrapperValueCodec is codec for Protobuf type wrappers
type wrapperValueCodec struct {
}

// EncodeValue encodes Protobuf type wrapper value to BSON value
func (e *wrapperValueCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	val = val.FieldByName("Value")
	enc, err := ectx.LookupEncoder(val.Type())
	if err != nil {
		return err
	}
	return enc.EncodeValue(ectx, vw, val)
}

// DecodeValue decodes BSON value to Protobuf type wrapper value
func (e *wrapperValueCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	val = val.FieldByName("Value")
	enc, err := ectx.LookupDecoder(val.Type())
	if err != nil {
		return err
	}
	return enc.DecodeValue(ectx, vr, val)
}

// timestampCodec is codec for Protobuf Timestamp
type timestampCodec struct {
}

// EncodeValue encodes Protobuf Timestamp value to BSON value
func (e *timestampCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
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
func (e *timestampCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
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

// objectIDCodec is codec for Protobuf ObjectId
type objectIDCodec struct {
}

// EncodeValue encodes Protobuf ObjectId value to BSON value
func (e *objectIDCodec) EncodeValue(ectx bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
	v := val.Interface().(mongodb.ObjectId)
	// Create primitive.ObjectId from string
	id, err := primitive.ObjectIDFromHex(v.Value)
	if err != nil {
		return err
	}
	enc, err := ectx.LookupEncoder(objectIDPrimitiveType)
	if err != nil {
		return err
	}
	return enc.EncodeValue(ectx, vw, reflect.ValueOf(id))
}

// DecodeValue decodes BSON value to ObjectId value
func (e *objectIDCodec) DecodeValue(ectx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	enc, err := ectx.LookupDecoder(objectIDPrimitiveType)
	if err != nil {
		return err
	}
	var id primitive.ObjectID
	if err = enc.DecodeValue(ectx, vr, reflect.ValueOf(&id).Elem()); err != nil {
		return err
	}
	oid := mongodb.ObjectId{Value: id.Hex()}
	if err != nil {
		return err
	}
	val.Set(reflect.ValueOf(oid))
	return nil
}

// Register registers Google protocol buffers types codecs
func Register(rb *bsoncodec.RegistryBuilder) *bsoncodec.RegistryBuilder {
	return rb.RegisterCodec(boolValueType, wrapperValueCodecRef).
		RegisterCodec(bytesValueType, wrapperValueCodecRef).
		RegisterCodec(doubleValueType, wrapperValueCodecRef).
		RegisterCodec(floatValueType, wrapperValueCodecRef).
		RegisterCodec(int32ValueType, wrapperValueCodecRef).
		RegisterCodec(int64ValueType, wrapperValueCodecRef).
		RegisterCodec(stringValueType, wrapperValueCodecRef).
		RegisterCodec(uint32ValueType, wrapperValueCodecRef).
		RegisterCodec(uint64ValueType, wrapperValueCodecRef).
		RegisterCodec(timestampType, timestampCodecRef).
		RegisterCodec(objectIDType, objectIDCodecRef)
}
