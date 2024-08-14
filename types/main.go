package types

type Todo struct {
	ID        string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}