package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		panic(err)
	}

	if _, err = db.Exec("CREATE TABLE foo (name VARCHAR, age INT);"); err != nil {
		panic(err)
	}
	fmt.Printf("CREATE OK\n")

	res, err := db.Exec(`INSERT INTO foo (name, age) VALUES (?, ?);`, "Alice", 40)
	if err != nil {
		panic(err)
	}
	n, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("INSERT: %d row(s) OK\n", n)

	rows, err := db.Query("SELECT name, age FROM foo WHERE name = ?", "Alice")
	if err != nil {
		panic(err)
	}
	var name string
	var age int
	for rows.Next() {
		if err := rows.Scan(&name, &age); err != nil {
			panic(err)
		}
		fmt.Printf("SELECT: %q %d\n", name, age)
	}

	res, err = db.Exec(`DELETE FROM foo WHERE name = ?`, "Alice")
	if err != nil {
		panic(err)
	}
	n, err = res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("DELETE: %d row(s) OK\n", n)
}
