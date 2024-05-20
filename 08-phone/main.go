package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"regexp"
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
	_, err = db.Exec("DELETE FROM phone_numbers")
	if err != nil {
		log.Fatalf("Failed to truncate table: %v", err)
	}

	stm, err := db.Prepare("INSERT INTO phone_numbers(phone) VALUES (?)")
	defer stm.Close()
	if err != nil {
		log.Fatal(err)
	}
	formatedPhone := regexp.MustCompile(`[-\s()]`)
	for _, phone := range phones {
		phone := formatedPhone.ReplaceAllString(phone, "")
		_, err := stm.Exec(phone)
		if err != nil {
			fmt.Println(err)
		}
	}

	rows, err := db.Query("SELECT phone from phone_numbers")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var listPhone []string
	for rows.Next() {
		var phone string
		err = rows.Scan(&phone)
		fmt.Println(phone)
		if err != nil {
			log.Fatal(err)
		}
		listPhone = append(listPhone, phone)
	}

}
