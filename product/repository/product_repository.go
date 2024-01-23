package repository

import (
	"context"
	"ekuator/product/config"
	"ekuator/product/models"
	"log"
)

type ProductInterface interface {
	InsertProduct(data models.Product) error
	UpdateProduct(id int, data models.Product) error
	GetAllProduct() ([]models.Product, error)
	DetailProduct(id int) (models.Product, error)
	DeleteProduct(id int) error
}

type ProductImpl struct {
	mysqlSession *config.MysqlSession
}

func NewProductRepository(cf *config.MysqlSession) ProductInterface {
	return &ProductImpl{
		mysqlSession: cf,
	}
}

func (c *ProductImpl) InsertProduct(data models.Product) error {
	query := `insert into Product (id, name, price, stock) values (?, ?, ?, ?)`
	insData, err := c.mysqlSession.Db.Prepare(query)
	if err != nil {
		log.Println("error database: ", err.Error())
		return err
	}

	insData.Exec(data.Id, data.Nama, data.Price, data.Stock)

	return nil
}
func (c *ProductImpl) UpdateProduct(id int, data models.Product) error {
	query := `UPDATE Product SET name=?, price=?, stock=? WHERE id=?`
	updData, err := c.mysqlSession.Db.Prepare(query)
	if err != nil {
		log.Println("error database : ", err.Error())
		return err
	}

	updData.Exec(data.Nama, data.Price, data.Stock, id)
	return nil
}

func (c *ProductImpl) GetAllProduct() ([]models.Product, error) {
	var (
		dataResult []models.Product
		product    models.Product
		err        error
	)

	dataResult = []models.Product{}

	query := `select * from Product`
	result, err := c.mysqlSession.Db.Query(query)
	if err != nil {
		return dataResult, err
	}

	for result.Next() {
		err := result.Scan(&product.Id, &product.Nama, &product.Price, &product.Stock, &product.CreatedDate, &product.UpdateDate)
		if err != nil {
			log.Println("error database: ", err.Error())
		}
		dataResult = append(dataResult, product)
	}

	return dataResult, nil
}

func (c *ProductImpl) DetailProduct(id int) (models.Product, error) {
	var (
		data models.Product
		err  error
	)
	data = models.Product{}

	query := `select * from Product where id=?`

	state, err := c.mysqlSession.Db.PrepareContext(context.TODO(), query)
	if err != nil {
		log.Println("error databse: ", err.Error())
		return data, err
	}
	err = state.QueryRowContext(context.TODO(), id).Scan(
		&data.Id,
		&data.Nama,
		&data.Price,
		&data.Stock,
		&data.CreatedDate,
		&data.UpdateDate,
	)
	if err != nil {
		log.Println("error databses: ", err.Error())
		return data, err
	}

	return data, nil
}

func (c *ProductImpl) DeleteProduct(id int) error {
	query := `DELETE FROM Product WHERE id=?`

	delCos, err := c.mysqlSession.Db.Prepare(query)
	if err != nil {
		return err
	}

	delCos.Exec(id)
	return nil
}
