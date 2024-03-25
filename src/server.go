package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"RestApiGo/src/db"
	"RestApiGo/src/routers"
)

/*
https://lucasfcosta.com/2017/02/07/Understanding-Go-Dependency-Management.html
https://gowebexamples.com/mysql-database/
https://go.dev/doc/effective_go
https://github.com/gorilla/mux?tab=readme-ov-file
https://www.youtube.com/watch?v=dbXgs-aQ7cE
*/

func main() {
	//https://github.com/gorilla/mux?tab=readme-ov-file
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello, you've requested: " + r.URL.Path))
	})
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	r.HandleFunc("/magazine/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		var params = mux.Vars(r)
		w.WriteHeader(200)
		w.Write([]byte("You requested books: " + params["title"] + " page " + params["page"]))
	})

	db.ConnectDb()

	routers.SetSubRouters(r)
	http.ListenAndServe(":8080", r)
}
