package entities

// Product presents a course product
type Product struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	ProducerId int64  `json:"producer_id"`
}
