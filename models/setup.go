package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectToDB() error {
	var err error
	DB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gogin_blog")
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	return nil
}
