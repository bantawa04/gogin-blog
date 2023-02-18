package setup

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectToDB() (*sql.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPORT := os.Getenv("DB_PORT")

	var err error
	//DB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/gogin_blog")
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPORT, dbName))

	if err != nil {
		return nil, err
	}
	return db, nil
}
