package entities

// DataEntry presents a data line in the data file
type DataEntry struct {
	Type    int    `json:"type"`
	Date    string `json:"date"`
	Product string `json:"product"`
	Value   int    `json:"value"`
	Seller  string `json:"seller"`
}
