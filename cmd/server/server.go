package main

import (
	"log"
	"net/http"

	"task31A/pkg/api"
	"task31A/pkg/storage"
	"task31A/pkg/storage/memdb"
	"task31A/pkg/storage/mongo"
	"task31A/pkg/storage/postgres"
)

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	// Создаём объект сервера.
	var srv server

	// Создаём объекты баз данных.
	//
	// БД в памяти.
	db1 := memdb.New()

	// Реляционная БД PostgreSQL.
	db2, err := postgres.New("postgres://postgres:password@server.domain:5432/posts")
	if err != nil {
		log.Fatal(err)
	}
	// Документная БД MongoDB.
	db3, err := mongo.New("mongodb://server.mongo:27017/")
	if err != nil {
		log.Fatal(err)
	}
	_, _, _ = db1, db2, db3

	// Инициализируем хранилище сервера конкретной БД.
	srv.db = db2

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	// Запускаем веб-сервер на порту 8080 на всех интерфейсах.
	// Предаём серверу маршрутизатор запросов,
	// поэтому сервер будет все запросы отправлять на маршрутизатор.
	// Маршрутизатор будет выбирать нужный обработчик.
	http.ListenAndServe(":8080", srv.api.Router())
}
