package database

import (
	"database/sql"

	"fmt"
    "time"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() sql.DB {
	// db, err := sql.Open("mysql", "root:@/golang")
	fmt.Println("Connecting to db")
	db, err := sql.Open("db", "root:mypassword@tcp(db:3306)/golang")
	for db.Ping() != nil {
		fmt.Println("Attempting connection to db")
		time.Sleep(5 * time.Second)
	}
	fmt.Println("Connected")
	checkErr(err)
	return *db
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Panic WOY")
		panic(err)
	}
}
