package services

import (
	"ekuator/costumer/models"
	"ekuator/costumer/repository"
	"log"
	"strconv"
)

type CostumerServiceInterface interface {
	SaveCostumer(data models.Costumer) error
	DeleteCostumer(id string) error
	DetailCostumer(id string) models.Costumer
	// UpdateCostumer() error
	GetAllCostumer() []models.Costumer
}

type CostumerServiceImpl struct {
	costumerRepo repository.CostumerInterface
}

func NewCostumerService(cr repository.CostumerInterface) CostumerServiceInterface {
	return &CostumerServiceImpl{
		costumerRepo: cr,
	}
}

func (cs *CostumerServiceImpl) SaveCostumer(data models.Costumer) error {
	err := cs.costumerRepo.InsertCostumer(data)
	if err != nil {
		log.Println("service insert costumer error: ", err.Error())
		return err
	}
	return nil
}

func (cs *CostumerServiceImpl) DeleteCostumer(id string) error {
	parseId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("error parse Id Costumer: ", err.Error())
		return err
	}

	err = cs.costumerRepo.DeleteCostumer(parseId)
	if err != nil {
		log.Println("Error delete Costumer : ", err.Error())
		return err
	}
	return nil
}

func (cs *CostumerServiceImpl) DetailCostumer(id string) models.Costumer {
	var (
		data models.Costumer
		err  error
	)

	parseId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("error parse string to int (id costumer): ", err.Error())
	}

	data, err = cs.costumerRepo.DetailCostumer(parseId)
	if err != nil {
		log.Println("error di service : ", err.Error())
		return models.Costumer{}
	}

	return data
}

// func (cs *CostumerServiceImpl) UpdateCostumer() error {}
func (cs *CostumerServiceImpl) GetAllCostumer() []models.Costumer {
	var (
		dataCostumer []models.Costumer
	)

	dataCostumer, err := cs.costumerRepo.GetAllCostumer()
	if err != nil {
		return dataCostumer
	}
	return dataCostumer
}
