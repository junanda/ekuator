package repository

import (
	"context"
	"ekuator/costumer/config"
	"ekuator/costumer/models"
	"log"
)

type CostumerInterface interface {
	InsertCostumer(data models.Costumer) error
	UpdateCostumer(id int, data models.Costumer) error
	GetAllCostumer() ([]models.Costumer, error)
	DetailCostumer(id int) (models.Costumer, error)
	DeleteCostumer(id int) error
}

type CostumerImpl struct {
	mysqlSession *config.MysqlSession
}

func NewCostumerRepository(cf *config.MysqlSession) CostumerInterface {
	return &CostumerImpl{
		mysqlSession: cf,
	}
}

func (c *CostumerImpl) InsertCostumer(data models.Costumer) error {
	query := `insert into Customer (id, name, email) values (?, ?, ?)`
	insData, err := c.mysqlSession.Db.Prepare(query)
	if err != nil {
		log.Println("error database: ", err.Error())
		return err
	}

	insData.Exec(data.Id, data.Nama, data.Email)

	return nil
}
func (c *CostumerImpl) UpdateCostumer(id int, data models.Costumer) error {
	query := `UPDATE Customer SET name=?, Email=? WHERE id=?`
	updData, err := c.mysqlSession.Db.Prepare(query)
	if err != nil {
		log.Println("error database : ", err.Error())
		return err
	}

	updData.Exec(data.Nama, data.Email, id)
	return nil
}

func (c *CostumerImpl) GetAllCostumer() ([]models.Costumer, error) {
	var (
		dataResult []models.Costumer
		customer   models.Costumer
		err        error
	)

	dataResult = []models.Costumer{}

	query := `select * from Customer`
	result, err := c.mysqlSession.Db.Query(query)
	if err != nil {
		return dataResult, err
	}

	for result.Next() {
		err := result.Scan(&customer.Id, &customer.Nama, &customer.Email, &customer.CreatedDate, &customer.UpdateDate)
		if err != nil {
			log.Println("error database: ", err.Error())
		}
		dataResult = append(dataResult, customer)
	}

	return dataResult, nil
}

func (c *CostumerImpl) DetailCostumer(id int) (models.Costumer, error) {
	var (
		data models.Costumer
		err  error
	)
	data = models.Costumer{}

	query := `select * from Customer where id=?`

	state, err := c.mysqlSession.Db.PrepareContext(context.TODO(), query)
	if err != nil {
		log.Println("error databse: ", err.Error())
		return data, err
	}
	err = state.QueryRowContext(context.TODO(), id).Scan(
		&data.Id,
		&data.Nama,
		&data.Email,
		&data.CreatedDate,
		&data.UpdateDate,
	)
	if err != nil {
		log.Println("error databses: ", err.Error())
		return data, err
	}

	return data, nil
}

func (c *CostumerImpl) DeleteCostumer(id int) error {
	query := `DELETE FROM Customer WHERE id=?`

	delCos, err := c.mysqlSession.Db.Prepare(query)
	if err != nil {
		return err
	}

	delCos.Exec(id)
	return nil
}
