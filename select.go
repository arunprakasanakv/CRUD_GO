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
	rows, err := db.Query("SELECT * FROM content")

	if err != nil {
		// handle this error better than this
		panic(err)
	}

	defer rows.Close()

	fmt.Println("id  |  Content")

	for rows.Next() {
		var id int
		var firstName string
		err = rows.Scan(&id, &firstName)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(id, "  |  ", firstName)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

}
