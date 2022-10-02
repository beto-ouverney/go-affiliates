package entities

// SaleResponse presents the response client of a sale
type SaleResponse struct {
	Product    string `json:"product" db:"product"`
	Producer   string `json:"producer" db:"producer"`
	Affiliate  string `json:"affiliate" db:"affiliate"`
	Value      int    `json:"value" db:"value"`
	Commission int    `json:"commission" db:"commission"`
	Date       string `json:"date" db:"date"`
}
