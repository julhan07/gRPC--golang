package entities

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	UserID   string  `json:"user_id"`
	UserName string  `json:"user_name"`
}
