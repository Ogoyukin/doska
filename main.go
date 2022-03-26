package main

import (
	"github.com/Ogoyukin/doska/pkg/db"
	"github.com/Ogoyukin/doska/pkg/handlers"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {

	// Устанавливаю соединение с Базой данных MySQL
	dataSource, err := db.ConnectDB()

	// Проверяю есть ли ошибки
	if err != nil {

		// Если есть ошибка, останавливаю приложение с ошибкой
		log.Fatalf("Не могу подключиться к БД: %v", err)
	}

	indexHandler := handlers.Handlers{
		DataSource: dataSource,
	}

	// Создаю обработчик запросов
	http.HandleFunc("/hello", indexHandler.Index)

	// Запускаю http server
	http.ListenAndServe(":8090", nil)
}
