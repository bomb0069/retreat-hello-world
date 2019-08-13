package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Say struct {
	Message string `db:"message"`
}

type SayRepository interface {
	GetFirstMessage() Say
}

type PostgresDB struct {
	Connection *sqlx.DB
}

func (db PostgresDB) GetFirstMessage() Say {
	var say Say
	err := db.Connection.Get(&say, "select * from say")
	if err != nil {
		log.Println("error ", err)
	}
	return say
}
