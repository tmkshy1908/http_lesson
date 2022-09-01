package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=yamadatarou dbname=test password=taisei3480 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

type Member struct {
	Id         int
	FirstName  string
	LastName   string
	Email      string
	AccessPrev bool
}

func main() {
	fmt.Println("----- Query 1 -----")
	rows, err := Db.Query("SELECT * FROM members ORDER BY id")
	if err != nil {
		return
	}
	for rows.Next() {
		m := Member{}
		rows.Scan(&m.Id, &m.FirstName, &m.LastName, &m.Email, &m.AccessPrev)
		fmt.Println(m)
	}
	fmt.Println()

	fmt.Println("----- Query 2 -----")
	rows, err = Db.Query("SELECT * FROM members WHERE accessprev = $1 ORDER BY id", "false")
	if err != nil {
		return
	}
	for rows.Next() {
		m := Member{}
		rows.Scan(&m.Id, &m.FirstName, &m.LastName, &m.Email, &m.AccessPrev)
		fmt.Println(m)
	}
	fmt.Println()

	fmt.Println("----- Prepare -----")
	statment := "SELECT * FROM members  WHERE accessprev = $1 ORDER BY id"
	stmt, err := Db.Prepare(statment)
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err = stmt.Query("false")
	for rows.Next() {
		m := Member{}
		rows.Scan(&m.Id, &m.FirstName, &m.LastName, &m.Email, &m.AccessPrev)
		fmt.Println(m)
	}
	fmt.Println()

	fmt.Println("----- Exec -----")
	_, err = Db.Exec("INSERT INTO members VALUES ($1, $2, $3, $4, $5)", 4, "takashi", "yamamoto", "yamamoto@yahoo.co.jp", "FALSE")
	if err != nil {
		return
	}
	rows, err = Db.Query("SELECT * FROM members ORDER BY id")
	if err != nil {
		return
	}
	for rows.Next() {
		m := Member{}
		rows.Scan(&m.Id, &m.FirstName, &m.LastName, &m.Email, &m.AccessPrev)
		fmt.Println(m)
	}
	fmt.Println()
}
