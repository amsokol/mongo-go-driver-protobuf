package mongodb

import (
	"encoding/json"
)

// MarshalJSON marshals ObjectId to JSON string
func (o ObjectId) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Id)
}

// UnmarshalJSON unmarshal JSON string to ObjectId
func (o *ObjectId) UnmarshalJSON(data []byte) error {
	var id string
	if err := json.Unmarshal(data, &id); err != nil {
		return err
	}
	o.Id = id
	return nil
}
