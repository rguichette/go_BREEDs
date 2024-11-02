package main

import (
	"flag"
	"fmt"
	"go_breeders/adapters"
	"go_breeders/config"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	App         *config.Application
}

type appConfig struct {
	useCache bool
	dsn      string
}

func main() {
	app := application{
		templateMap: make(map[string]*template.Template),
	}
	flag.BoolVar(&app.config.useCache, "cache", false, "use template cache")
	flag.StringVar(&app.config.dsn, "dsn", "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "DSN")

	//connect to db

	db, err := initMySQLDB(app.config.dsn)
	if err != nil {
		log.Panic(err)
	}

	// JsonBackend := &adapters.JSONBackend{}
	// jsonAdapter := &adapter.RemoteService{Remote: JsonBackend}
	xmlBackend := &adapters.XMLBackEnd{}
	xmlAdapter := &adapters.RemoteService{Remote: xmlBackend}

	app.App = config.New(db, xmlAdapter)
	// app.catService = jsonAdapter
	flag.Parse()

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("started web app on port...", port)

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
