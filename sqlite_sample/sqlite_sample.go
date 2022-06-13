package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./testdb")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	//tableName := "testtable"
	//table_create(db, tableName)

	// cmd := "INSERT INTO testtable (id, name, age, email) VALUES ($1, $2, $3, $4)"
	// result, err := db.Exec(cmd, 1, "Taro", 20, "xxxxxx@yyyyyy.com")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// fmt.Println(result)

	cmd := "SELECT * FROM testtable"
	result, err := db.Query(cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer result.Close()
	cullum, err := result.Columns()
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	var name, email string
	var id, age int

	fmt.Println(cullum)
	result.Next()
	err = result.Scan(&id, &name, &age, &email)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println(id, name, email, age)
}

func insert_() {

}

func table_create(db *sql.DB, tableName string) {
	cmd := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id SERIAL PRIMARY KEY NOT NULL,
		name TEXT NOT NULL,
		age INTEGER NOT NULL,
		email TEXT)`, tableName)
	_, err := db.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}
}
