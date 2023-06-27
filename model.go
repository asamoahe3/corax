package main

type DummyProducts struct {
	Products []struct {
		ID                 int      `json:"id"`
		Title              string   `json:"title"`
		Description        string   `json:"description"`
		Price              int      `json:"price"`
		DiscountPercentage float64  `json:"discountPercentage"`
		Rating             float64  `json:"rating"`
		Stock              int      `json:"stock"`
		Brand              string   `json:"brand"`
		Category           string   `json:"category"`
		Thumbnail          string   `json:"thumbnail"`
		Images             []string `json:"images"`
	} `json:"products"`
	Total int `json:"total"`
	Skip  int `json:"skip"`
	Limit int `json:"limit"`
}
