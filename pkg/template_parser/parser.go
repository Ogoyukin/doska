package template_parser

import "html/template"

func Parse() (*template.Template, error) {

	// Подключаю корневой шаблон страницы
	tmp, err := template.ParseFiles("templates/index.html")

	if err != nil {
		return nil, err
	}

	// Сканирую другие шаблоны
	tmp, err = tmp.ParseGlob("templates/*.html")

	if err != nil {
		return nil, err
	}

	return tmp, err
}
