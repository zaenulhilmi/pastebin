package main

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/zaenulhilmi/pastebin/handlers"
	"github.com/zaenulhilmi/pastebin/helpers"
	"github.com/zaenulhilmi/pastebin/repositories"
	"github.com/zaenulhilmi/pastebin/services"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", handlers.PingHandler)

	db := InitDB()
	clock := helpers.SystemClock{}
	readPasteRepository := repositories.NewReadPasteRepository(db, clock)
	writePasteRepository := repositories.NewWritePasteRepository(db, clock)

	cache := helpers.NewCache()
	readCacheRepositoryAdapter := repositories.NewCacheAdapter(readPasteRepository, cache)
	shortlinkGenerator := services.NewShortlinkGenerator(readCacheRepositoryAdapter, &helpers.DefaultToken{})

	pasteService := services.NewShortlinkService(readPasteRepository, writePasteRepository, shortlinkGenerator)
	pasteHandler := handlers.NewShortlinkHandler(pasteService)

	logRepository := repositories.NewLogRepository(db)
	logService := services.NewLogService(logRepository)
	pasteHandle := handlers.LoggingMiddleware(logService, http.HandlerFunc(pasteHandler.Content))

	r.Handle("/paste", pasteHandle)

	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Minutes().Do(func() {
		pasteService.DeleteExpiredContent()
	})
	s.StartAsync()

	http.ListenAndServe(":8080", r)
}

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "groot:password@tcp(localhost:3306)/pastebin?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}
