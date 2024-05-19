package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var phones = []string{
	"1234567890",
	"123 456 7891",
	"(123) 456 7892",
	"(123) 456-7893",
	"123-456-7894",
	"123-456-7890",
	"1234567892",
	"(123)456-7892",
}

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	createTableSQL := `CREATE TABLE IF NOT EXISTS phone_numbers (
        "id" INTEGER PRIMARY KEY AUTOINCREMENT,
        "phone" TEXT
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Table created successfully!")
}
