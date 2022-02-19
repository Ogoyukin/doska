package main

import (
	"database/sql"
	"fmt"
	"github.com/Ogoyukin/doska/pkg"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

var db *sql.DB

func main() {
	dataSource, err := sql.Open("mysql", "root:root@/doska")

	db = dataSource
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/hello", Handler)
	http.ListenAndServe(":8090", nil)
}

type Params struct {
	Counter int
	Items   []pkg.Item
}

func Handler(w http.ResponseWriter, r *http.Request) {

	tmp, _ := template.ParseFiles("templates/index.html")
	tmp.ParseGlob("templates/*.html")

	products := []pkg.Item{}

	rows, err := db.Query("select * from doska.items")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		p := pkg.Item{}
		err := rows.Scan(&p.ID, &p.Price, &p.Title, &p.Description, &p.Img)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	tmp.Execute(w, Params{Items: products})
}
