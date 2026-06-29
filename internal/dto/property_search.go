package dto

type PropertySearchRequest struct {
    City            string  `form:"city"`
    TransactionType string  `form:"transaction_type"`
    Status          string  `form:"status"`
    MinPrice        float64 `form:"min_price"`
    MaxPrice        float64 `form:"max_price"`
    Bedrooms        int     `form:"bedrooms"`
    Bathrooms       int     `form:"bathrooms"`
    Page            int     `form:"page"`
    Limit           int     `form:"limit"`
    Sort            string  `form:"sort"`
}