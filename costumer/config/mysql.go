package config

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlSession struct {
	Db *sql.DB
}

func NewMysqlDatabase() *MysqlSession {
	mysql := new(MysqlSession)
	return mysql
}

func (m *MysqlSession) Initialize() *MysqlSession {
	var err error
	m.Db, err = sql.Open("mysql", "user:password@/ekuator_test")
	if err != nil {
		log.Printf("Unable to Connect DB to %v", err)
		return nil
	}

	m.Db.SetMaxIdleConns(10)
	m.Db.SetConnMaxIdleTime(time.Duration(10))
	err = m.Db.Ping()
	if err != nil {
		log.Printf("Unable to Connect DB to %v", err)
		return nil
	}

	log.Printf("Connected to MySQL %s!\n", "ekuator_test")
	return m
}
