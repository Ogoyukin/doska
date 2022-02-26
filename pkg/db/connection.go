package db

import "database/sql"

// Устанавливаю соединение с Базой данных MySQL
func ConnectDB() (*sql.DB, error) {
	dataSource, err := sql.Open("mysql", "root:root@/doska")

	if err != nil {
		return nil, err
	}

	// Проверяю соединение
	err = dataSource.Ping()

	if err != nil {
		return nil, err
	}

	return dataSource, err
}
