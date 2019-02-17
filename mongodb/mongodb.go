package mongodb

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// GetObjectID returns MongoDB object ID
func (o *ObjectId) GetObjectID() (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(o.Value)
}
