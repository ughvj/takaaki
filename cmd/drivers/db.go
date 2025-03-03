package drivers

import (
	"database/sql"
)

type DBConnector interface {
	Use() *sql.DB
}
