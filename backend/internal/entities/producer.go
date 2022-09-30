package entities

// Producer presents a content content producer
// ID content producer id in database
// Name content producer name
type Producer struct {
	ID   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
