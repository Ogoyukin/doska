package main

import (
	"github.com/Ogoyukin/doska/pkg"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", Handler)
	http.ListenAndServe(":8090", nil)
}

var counter = 0

type Params struct {
	Counter int
	Items   []pkg.Item
}

func Handler(w http.ResponseWriter, r *http.Request) {

	counter++
	tmp, _ := template.ParseFiles("templates/index.html")
	tmp.ParseGlob("templates/*.html")

	items := []pkg.Item{{
		ID:          1,
		Price:       50000,
		Title:       " Продаётся iPhone 13",
		Description: "Продаётся iPhone 13 Pro пользовался только неделю, в отл состояниее акум 99%. Только сегодня...",
		Img:         "https://img.mvideo.ru/Pdb/30059037b.jpg",
	}}

	for i := 0; i < 20; i++ {
		img := "https://telefonplus.ru/wp-content/uploads/2020/10/iphone-12-blue-select-2020.png"

		if i%2 == 0 {
			img = "https://img.mvideo.ru/Pdb/30059037b.jpg"
		}
		items = append(items, pkg.Item{
			ID:          1,
			Price:       50000,
			Title:       " Продаётся iPhone 12",
			Description: "Продаётся iPhone 12  пользовался только неделю, в отл состояниее акум 99%. Только сегодня...",
			Img:         img,
		})
	}

	tmp.Execute(w, Params{Items: items})
}
