package entities

// Sale represents a sale
type Sale struct {
	ID          int64  `db:"id"`
	producerId  int64  `db:"producer_id"`
	affiliateId int64  `db:"affiliate_id"`
	productId   int64  `db:"product_id"`
	value       int    `db:"value"`
	date        string `db:"date"`
}
