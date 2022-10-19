package tools

import (
	"fmt"

	"github.com/PedroSMarcal/todo_list_golang/constants"
)

func SetPort(port *string) {
	*port = fmt.Sprintf(":%s", *port)
}

func CreateDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		constants.TodoDatabaseHost,
		constants.TodoDatabaseUser,
		constants.TodoDatabasePassword,
		constants.TodoDatabaseDatabase,
		constants.TodoDatabasePort,
	)
}
