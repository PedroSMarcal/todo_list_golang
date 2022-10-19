package tools

import (
	"fmt"

	"github.com/PedroSMarcal/todo_list_golang/config"
)

func SetPort(port *string) {
	*port = fmt.Sprintf(":%s", *port)
}

func CreateDSN() string {

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		config.EnvVariable.DbHost, config.EnvVariable.DbUser, config.EnvVariable.DbPassword, config.EnvVariable.DbName, config.EnvVariable.DbPort,
	)
}
