package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	
	"database/sql"
	"html/template"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	_"github.com/lib/pq"
)

var tmplPath = "./templates/"

var db *sql.DB

type Tweets struct {
	Url        string
	Text       string
	Tweeted_at time.Time
}

func handler_index(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(tmplPath + "index.html")
	if err != nil {
		log.Fatal(err)
	}

	var sql = "select url, text, tweeted_at from tweets"

	search, ok := r.URL.Query()["search"]
	if ok {
		sql += " where text like '%" + strings.Replace(search[0], "'", "\\'", -1) + "%'"
	}

	sql += " order by tweeted_at desc"

	limit, ok := r.URL.Query()["limit"]
	if ok && (limit[0] != "") {
		sql += " limit " + strings.Split(limit[0], ";")[0]
	}

	var data []Tweets


	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, sql)
	if err != nil{
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for rows.Next() {
		var text string
		var url string
		var tweeted_at time.Time

		err := rows.Scan(&url, &text, &tweeted_at)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		data = append(data, Tweets{url, text, tweeted_at})
	}

	tmpl.Execute(w, data)
}

func initialize() {
	var err error

	dbname := "ctf"
	dbuser := os.Getenv("FLAG")
	dbpass := "password"

	connInfo := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable", 5432, "db", dbuser, dbpass, dbname)
	db, err = sql.Open("postgres", connInfo)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	initialize()

	r := mux.NewRouter()
	r.HandleFunc("/", handler_index).Methods("GET")

	http.Handle("/", r)
	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
