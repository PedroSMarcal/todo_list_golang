package tools

import "fmt"

func SetPort(port *string) {
	*port = fmt.Sprintf(":%s", *port)
}
