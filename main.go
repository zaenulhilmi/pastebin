package main

import (
	"fmt"
    "github.com/gorilla/mux"
    "github.com/zaenulhilmi/pastebin/handlers"
    "github.com/zaenulhilmi/pastebin/services"
    "github.com/zaenulhilmi/pastebin/repositories"
    _ "github.com/go-sql-driver/mysql"
    "net/http"
    "database/sql"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/ping", handlers.PingHandler)
    
    db := InitDB()
    pasteRepository := repositories.NewShortlinkRepository(db)
    pasteService := services.NewShortlinkService(pasteRepository)
    pasteHandler := handlers.NewShortlinkHandler(pasteService)

    r.HandleFunc("/paste", pasteHandler.GetContent)

    http.ListenAndServe(":8080", r)

}

func InitDB() *sql.DB {
    db, err := sql.Open("mysql", "groot:password@tcp(localhost:3306)/pastebin?parseTime=true")
    if err != nil {
        panic(err.Error())
    }
    return db
}
