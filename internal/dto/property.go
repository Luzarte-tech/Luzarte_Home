package dto

type CreatePropertyRequest struct {
	CategoryID      string  `json:"category_id" binding:"required"`
	Title           string  `json:"title" binding:"required"`
	Description     string  `json:"description"`
	TransactionType string  `json:"transaction_type" binding:"required"`
	Price           float64 `json:"price" binding:"required"`
	Bedrooms        int     `json:"bedrooms"`
	Bathrooms       int     `json:"bathrooms"`
	GarageSpaces    int     `json:"garage_spaces"`
	Area            float64 `json:"area"`
	Address         string  `json:"address"`
	City            string  `json:"city"`
	Province        string  `json:"province"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Featured        bool    `json:"featured"`
	Published       bool    `json:"published"`
}

type UpdatePropertyRequest struct {
	CategoryID      string  `json:"category_id"`
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	TransactionType string  `json:"transaction_type"`
	Price           float64 `json:"price"`
	Bedrooms        int     `json:"bedrooms"`
	Bathrooms       int     `json:"bathrooms"`
	GarageSpaces    int     `json:"garage_spaces"`
	Area            float64 `json:"area"`
	Address         string  `json:"address"`
	City            string  `json:"city"`
	Province        string  `json:"province"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Status          string  `json:"status"`
	Featured        bool    `json:"featured"`
	Published       bool    `json:"published"`
}