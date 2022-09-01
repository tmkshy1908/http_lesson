package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)

// score management

type Students struct {
	Id      int
	Name    string
	Lang    int
	Re_lang int
	Math    int
	Re_math int
	Hist    int
	Re_hist int
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=yamadatarou dbname=students password=taisei3480 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func seting() (err error) {
	var a [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 5; i++ {
		a[i-1] = rand.Intn(100)
		students := &Students{Id: i, Lang: a[i-1]}
		_, err = Db.Exec("update students set lang = $1 where id = $2", students.Lang, students.Id)
	}
	for i := 1; i <= 5; i++ {
		a[i-1] = rand.Intn(100)
		students := &Students{Id: i, Math: a[i-1]}
		_, err = Db.Exec("update students set math = $1 where id = $2", students.Math, students.Id)
	}
	for i := 1; i <= 5; i++ {
		a[i-1] = rand.Intn(100)
		students := &Students{Id: i, Hist: a[i-1]}
		_, err = Db.Exec("update students set Hist = $1 where id = $2", students.Hist, students.Id)
	}
	return
}

func avg() {
	// s := Students{}
	lang_av, err := Db.Query("select avg(lang) from students")
	// lang_av.Scan(s.Lang)
	fmt.Println(lang_av)
	// math_av, err := Db.Query("select avg(math) from students")
	// hist_av, err := Db.Query("select avg(hist) from students")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(lang_av)
}

func main() {
	seting()
	rows, err := Db.Query("select * from students")
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		s := &Students{}
		if err := rows.Scan(&s.Id, &s.Name, &s.Lang, &s.Re_lang, &s.Math, &s.Re_math, &s.Hist, &s.Re_hist); err != nil {
			fmt.Println(err)
		}
		fmt.Println(s)
	}
	avg()
}
