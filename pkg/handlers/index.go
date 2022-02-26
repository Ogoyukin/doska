package handlers

import (
	"database/sql"
	"fmt"
	"github.com/Ogoyukin/doska/pkg"
	"github.com/Ogoyukin/doska/pkg/template_parser"
	"log"
	"net/http"
)

type Handlers struct {
	DataSource *sql.DB
}

func (h *Handlers) Index(w http.ResponseWriter, r *http.Request) {

	tmp, err := template_parser.Parse()

	// Проверяю есть ли ошибки
	if err != nil {

		// Если есть ошибка, останавливаю приложение с ошибкой
		log.Fatalf("Не могу подключиться к БД: %v", err)
	}

	items := []pkg.Item{}

	// Отправляю запрос в БД, достать данные из таблицы items в базе doska
	rows, err := h.DataSource.Query("select * from doska.items")

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

type Params struct {
	Items []pkg.Item
}
