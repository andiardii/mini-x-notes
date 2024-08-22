package core

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitializeDatabase() {
    dsn := "root:123@tcp(mysql:3306)/db"
    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        panic("Failed to connect to database: " + err.Error())
    }
}
