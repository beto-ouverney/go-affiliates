package entities

// Affiliate presents a affiliate content producer
type Affiliate struct {
	ID         int64  `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	ProducerId int64  `json:"producer_id" db:"producer_id"`
}
