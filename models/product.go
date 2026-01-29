package models

type Product struct {
    ID          int     `json:"id"`
    Name        string  `json:"name"`
    Price       float64 `json:"price"`
    CategoryID  int     `json:"category_id"` // Ini foreign key-nya
}