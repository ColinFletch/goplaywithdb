package driver

import (
	"database/sql"
	"fmt"
)

// DB
type DB struct {
	SQL *sql.DB
}

//DB Connection
var dbConn = &DB{}

//connect
func ConnectSQL(host, port, uname, pass, dbname string) (*DB, error) {
	dbSource := fmt.Sprintf(
		"root:%s@tcp(%s:%s)/%s?charset=utf8",
		pass,
		host,
		port,
		dbname,
	)
	d, err := sql.Open("mysql", dbSource)
	if err != nil { // freak out and log error connecting to DB
		panic(err)
	}
	dbConn.SQL = d
	return dbConn, err
}
