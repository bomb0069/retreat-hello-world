package repository

import (
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Test_GetFirstMessage_Should_Be_Hello_World(t *testing.T) {
	expectedModel := Say{Message: "Hello World"}

	dataSource := fmt.Sprintf("host=localhost port=5432 user=postgres password=mysecretpassword dbname=retreat sslmode=disable")
	connection, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		t.Errorf("Connection Error : %s", err.Error())
	}

	postgresDB := PostgresDB{Connection: connection}

	actualModel := postgresDB.GetFirstMessage()
	if actualModel != expectedModel {
		t.Errorf("actual %+v not equals to expected %+v", actualModel, expectedModel)
	}

}
