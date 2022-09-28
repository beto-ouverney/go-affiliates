package entities

// Affiliate presents a affiliate content producer
type Affiliate struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	ProducerId int64  `json:"producer_id"`
}
