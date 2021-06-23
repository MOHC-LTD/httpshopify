package httpshopify

type TaxLinesDTO []TaxLineDTO

type TaxLineDTO struct {
	Title string  `json:"title"`
	Rate  float32 `json:"rate"`
}
