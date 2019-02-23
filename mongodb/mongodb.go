package mongodb

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetObjectID returns MongoDB object ID
func (o *ObjectId) GetObjectID() (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(o.Value)
}
