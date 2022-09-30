package entities

// Product presents a course product
type Product struct {
	ID         int64  `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	ProducerId int64  `json:"producer_id" db:"producer_id"`
}
