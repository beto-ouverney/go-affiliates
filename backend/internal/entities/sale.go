package entities

// Sale represents a sale
// ID sale id in database
// ProductId content producer id of this product
// AffiliateId affiliate id in database
// ProducerId content producer id of this product
// Value sale value
// Commission sale commission of the affiliate
// Date sale date
type Sale struct {
	ID          int64  `db:"id"`
	ProducerId  int64  `db:"producer_id"`
	AffiliateId int64  `db:"affiliate_id"`
	ProductId   int64  `db:"product_id"`
	Value       int    `db:"value"`
	Commission  int    `db:"commission"`
	Date        string `db:"date"`
}
