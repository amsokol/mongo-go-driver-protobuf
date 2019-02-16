package mongodb

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// GetPrimitiveObjectID returns MongoDB object ID
func (o *ObjectId) GetPrimitiveObjectID() (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(o.Value)
}
