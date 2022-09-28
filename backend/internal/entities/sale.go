package entities

// Sale represents a sale
type Sale struct {
	ID          int64  `db:"id"`
	ProducerId  int64  `db:"producer_id"`
	AffiliateId int64  `db:"affiliate_id"`
	ProductId   int64  `db:"product_id"`
	Value       int    `db:"value"`
	Commission  int    `db:"commission"`
	Date        string `db:"date"`
}
