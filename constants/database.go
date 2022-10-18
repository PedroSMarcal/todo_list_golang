package constants

import "os"

const (
	DatabaseType string = `mysql`
)

var (
	DSN string = os.Getenv("DATABASE_URL")
)
