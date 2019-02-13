package mongodb

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// GetObjectId returns MongoDB object ID
func (o *ObjectId) GetObjectId() (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(o.Value)
}
