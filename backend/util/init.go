package util

import (
	"database/sql"

	"github.com/Githaiga22/forum/backend/DB"
)

var Database *sql.DB

func Init() {
	Database = DB.CreateConnection()
}
