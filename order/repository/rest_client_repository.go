package repository

import "ekuator/order/models"

type RestClientRepository interface {
	GetCostumer(id string) (models.Costumer, error)
	GetProduct(id string) (models.Product, error)
}

type RestClient struct{}

func (c *RestClient) GetCostumer(id string) (models.Costumer, error) {}

func (c *RestClient) GetProduct(id string) (models.Product, error) {}

func (c *RestClient) getRequest(url string) interface{} {

}
