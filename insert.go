package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	"log"
)

func main() {

	db, err := sql.Open("postgres", "user=postgres password='test'  dbname=test sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	sqlStatement := `
      INSERT INTO content (content)  
      VALUES ('asdasfsdfdsfsadfsdfgdsfgdsg')  
      RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)

}
