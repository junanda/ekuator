package models

type Order struct {
	Id         int     `json:"id"`
	CostumerId int     `json:"customer_id"`
	ProductId  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	Total      float64 `json:"total"`
	CreatedAt  string  `json:"created_at"`
	UpdateAt   string  `json:"update_at"`
}

type DetailOrder struct {
	Id            int     `json:"id"`
	NamaCustomer  string  `json:"nama_customer"`
	EmailCustomer string  `json:"email_customer"`
	NamaProduct   string  `json:"nama_product"`
	Quantity      int     `json:"quantity"`
	Total         float64 `json:"total"`
}
