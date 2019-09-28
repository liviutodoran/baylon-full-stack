package models

type Minimumwage struct {
	Id              interface{} `bson:"_id,omitempty"`
	Country     string
	Year        string
	LocalAmount string
	USD         string
}
