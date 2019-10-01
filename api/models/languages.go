package models

type Languages struct {
	ID   interface{} `bson:"_id,omitempty"`
	Name string
	Code string
}
