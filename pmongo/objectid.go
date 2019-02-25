package pmongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// NewObjectId creates proto ObjectId from MongoDB ObjectID
func NewObjectId(id primitive.ObjectID) *ObjectId {
	return &ObjectId{Value: id.Hex()}
}

// GetObjectID returns MongoDB object ID
func (o *ObjectId) GetObjectID() (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(o.Value)
}
