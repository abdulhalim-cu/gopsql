package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func main() {
	arguments := os.Args
	if len(arguments) != 6 {
		fmt.Println("Please provide: hostname port username password db")
		return
	}

	host := arguments[1]
	p := arguments[2]
	username := arguments[3]
	password := arguments[4]
	database := arguments[5]

	port, err := strconv.Atoi(p)
	if err != nil {
		fmt.Println("Not a valid port number: ", err)
		return
	}
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, database)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Println("Open():", err)
		return
	}
	defer db.Close()

	rows, err := db.Query(`SELECT "name", "amount_total" FROM "sale_order"`)
	if err != nil {
		fmt.Println("Query", err)
		return
	}
	for rows.Next() {
		var name string
		var total float64
		err = rows.Scan(&name, &total)
		if err != nil {
			fmt.Println("Scan", err)
			return
		}
		fmt.Println("*", name, total)
	}
	defer rows.Close()

	// get all table names from the __current__ database
	query := `SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' ORDER BY table_name`
	rows, err = db.Query(query)
	if err != nil {
		fmt.Println("Query", err)
		return
	}
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Println("Scan", err)
			return
		}
		fmt.Println("+T", name)
	}
	defer rows.Close()
}
