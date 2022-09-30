package entities

// Producer presents a content producer
type Producer struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
