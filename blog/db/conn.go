package db

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/blog?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
