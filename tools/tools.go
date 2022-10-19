package tools

import (
	"fmt"

	"github.com/PedroSMarcal/todo_list_golang/config"
)

func SetPort(port *string) {
	*port = fmt.Sprintf(":%s", *port)
}

func CreateDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.TodoDatabase.Host,
		config.TodoDatabase.User,
		config.TodoDatabase.Password,
		config.TodoDatabase.Database,
		config.TodoDatabase.Port,
	)
}
