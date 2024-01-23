package models

type Product struct {
	Id          int    `json:"id"`
	Nama        string `json:"name"`
	Price       string `json:"price"`
	Stock       string `json:"stock"`
	CreatedDate string `json:"created_at,omitempty"`
	UpdateDate  string `json:"update_at,omitempty"`
}
