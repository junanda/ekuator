package models

type Costumer struct {
	Id          int    `json:"id"`
	Nama        string `json:"name"`
	Email       string `json:"email"`
	CreatedDate string `json:"created_at,omitempty"`
	UpdateDate  string `json:"update_at,omitempty"`
}
