package entities

// DataEntry presents a data line in the data file
type DataEntry struct {
	Type    string `json:"type"`
	Data    string `json:"data"`
	Product string `json:"product"`
	Value   int    `json:"value"`
	Seller  string `json:"seller"`
}
