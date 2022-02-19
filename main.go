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

	// Устанавливаю соединениес Базой данных MySQL
	db, err := sql.Open("mysql", "root:root@/doska")

	// Проверяю есть ли ошибки
	if err != nil {

		// Если есть ошибка, останавливаю приложение с ошибкой
		log.Fatalf("Не могу подключиться к БД: %v", err)
	}

	// Проверяю соединение
	err = db.Ping()

	// Проверяю есть ли ошибки
	if err != nil {
		// Если есть ошибка, останавливаю приложение с ошибкой
		log.Fatalf("Не могу дойти до БД: %v", err)
	}

	// Создаю обработчик запросов
	http.HandleFunc("/hello", Handler)

	// Запускаю http server
	http.ListenAndServe(":8090", nil)
}

type Params struct {
	Counter int
	Items   []pkg.Item
}

func Handler(w http.ResponseWriter, r *http.Request) {

	// Подключаю корневой шаблон страницы
	tmp, _ := template.ParseFiles("templates/index.html")

	// Сканирую другие шаблоны
	tmp, err := tmp.ParseGlob("templates/*.html")

	items := []pkg.Item{}

	// Отправляю запрос в БД, достать данные из таблицы items в базе doska
	rows, err := db.Query("select * from doska.items")

	// Проверяю результат запроса на предмет ошибкок
	if err != nil {
		// Если есть ошибка, останавливаю приложение с ошибкой
		log.Fatalf("Не могу выполнить запрос: %v", err)
	}

	// Закрываю запрос, как только функция выполнится, ключевое слово "defer"
	defer rows.Close()

	// Скаинрую ответ от БД
	for rows.Next() {
		p := pkg.Item{}

		// Присваию ответы в модели item
		err := rows.Scan(&p.ID, &p.Price, &p.Title, &p.Description, &p.Img)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Добавляю элемент в массив
		items = append(items, p)
	}

	// Передаю шаблону необходимые параметры
	tmp.Execute(w, Params{Items: items})
}
