package entities

// Product presents a course product
// ID product id in database
// Name product name
// ProducerId  content producer id of this product
type Product struct {
	ID         int64  `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	ProducerId int64  `json:"producer_id" db:"producer_id"`
}
