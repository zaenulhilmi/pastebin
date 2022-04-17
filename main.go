package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/zaenulhilmi/pastebin/handlers"
	"github.com/zaenulhilmi/pastebin/helpers"
	"github.com/zaenulhilmi/pastebin/repositories"
	"github.com/zaenulhilmi/pastebin/services"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", handlers.PingHandler)

	db := InitDB()
	clock := helpers.SystemClock{}
	pasteRepository := repositories.NewShortlinkRepository(db, clock)

	cache := helpers.NewCache()
	cacheRepositoryAdapter := repositories.NewCacheAdapter(pasteRepository, cache)
	shortlinkGenerator := services.NewShortlinkGenerator(cacheRepositoryAdapter, &helpers.DefaultToken{})

	pasteService := services.NewShortlinkService(cacheRepositoryAdapter, shortlinkGenerator)
	pasteHandler := handlers.NewShortlinkHandler(pasteService)

	r.HandleFunc("/paste", pasteHandler.GetContent)
	r.HandleFunc("/create-paste", pasteHandler.CreateContent)

	http.ListenAndServe(":8080", r)

}

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "groot:password@tcp(localhost:3306)/pastebin?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}
