package repository

import (
	"context"
	"ekuator/order/config"
	"ekuator/order/models"
)

type OrderRepository interface {
	InsertOrder(data models.Order) error
	DetailOrder(id int) (models.Order, error)
	DeleteOrder(id string) error
}

type orderRepositoryImpl struct {
	mysqlDB *config.MysqlSession
}

func NewOrderRepository(cf *config.MysqlSession) OrderRepository {
	return &orderRepositoryImpl{
		mysqlDB: cf,
	}
}

func (o *orderRepositoryImpl) InsertOrder(data models.Order) error {
	query := `INSERT INTO Order (id, customer_id, product_id, quantity, total) VALUES (?, ?, ?, ?, ?)`
	instData, err := o.mysqlDB.Db.Prepare(query)
	if err != nil {
		return err
	}

	instData.Exec(data.Id, data.CostumerId, data.ProductId, data.Quantity, data.Total)
	return nil
}

func (o *orderRepositoryImpl) DetailOrder(id int) (models.Order, error) {
	var (
		data models.Order
		err  error
	)
	query := `SELECT * FROM Order WHERE id=?`

	getStat, err := o.mysqlDB.Db.PrepareContext(context.TODO(), query)
	if err != nil {
		return data, err
	}

	err = getStat.QueryRowContext(context.TODO()).Scan(
		&data.Id,
		&data.CostumerId,
		&data.ProductId,
		&data.Quantity,
		&data.Total,
		&data.CreatedAt,
		&data.UpdateAt,
	)

	if err != nil {
		return data, err
	}

	return data, nil
}

func (o *orderRepositoryImpl) DeleteOrder(id string) error {
	query := `DELETE FROM Order WHERE id=?`

	delStat, err := o.mysqlDB.Db.Prepare(query)
	if err != nil {
		return err
	}

	delStat.Exec(id)

	return nil
}
