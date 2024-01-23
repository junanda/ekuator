package services

import (
	"ekuator/product/models"
	"ekuator/product/repository"
	"log"
	"strconv"
)

type ProductServiceInterface interface {
	SaveProduct(data models.Product) error
	DeleteProduct(id string) error
	DetailProduct(id string) models.Product
	// UpdateProduct() error
	GetAllProduct() []models.Product
}

type ProductServiceImpl struct {
	prodRepo repository.ProductInterface
}

func NewProductService(cr repository.ProductInterface) ProductServiceInterface {
	return &ProductServiceImpl{
		prodRepo: cr,
	}
}

func (cs *ProductServiceImpl) SaveProduct(data models.Product) error {
	err := cs.prodRepo.InsertProduct(data)
	if err != nil {
		log.Println("service insert costumer error: ", err.Error())
		return err
	}
	return nil
}

func (cs *ProductServiceImpl) DeleteProduct(id string) error {
	parseId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("error parse Id Costumer: ", err.Error())
		return err
	}

	err = cs.prodRepo.DeleteProduct(parseId)
	if err != nil {
		log.Println("Error delete Costumer : ", err.Error())
		return err
	}
	return nil

}
func (cs *ProductServiceImpl) DetailProduct(id string) models.Product {
	var (
		data models.Product
		err  error
	)

	parseId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("error parse string to int (id costumer): ", err.Error())
	}

	data, err = cs.prodRepo.DetailProduct(parseId)
	if err != nil {
		log.Println("error di service : ", err.Error())
		return models.Product{}
	}

	return data
}

// func (cs *CostumerServiceImpl) UpdateCostumer() error {}
func (cs *ProductServiceImpl) GetAllProduct() []models.Product {
	var (
		dataCostumer []models.Product
	)

	dataCostumer, err := cs.prodRepo.GetAllProduct()
	if err != nil {
		return dataCostumer
	}
	return dataCostumer
}
