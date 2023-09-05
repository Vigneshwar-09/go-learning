package model

type Todo struct {
	ID     string `json:"id" bson:"_id,omitempty"`
	Status bool      `bson:"status"`
	Task   string    `bson:"task"`
}
