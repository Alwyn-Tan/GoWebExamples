package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

func main() {
	db, open := sql.Open("mysql", "root:1q2w3e4r5t@(127.0.0.1:3306)/test?parseTime=true")

	if open != nil {
		fmt.Fprintf(os.Stderr, open.Error())
	}
	err := db.Ping()

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}

	{
		query := `CREATE TABLE IF NOT EXISTS users (
    	id INT AUTO_INCREMENT PRIMARY KEY,
    	username TEXT NOT NULL,
    	password TEXT NOT NULL,
   		created_at DATETIME);`

		if _, err := db.Exec(query); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}
	}

	{
		username := "Irene"
		password := "123456"
		createdAt := time.Now()

		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?,?,?)`,
			username, password, createdAt)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
		}

		id, _ := result.LastInsertId()
		fmt.Println(id)
	}

	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)

	query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
	if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
	fmt.Println(id, username, password, createdAt)
}
