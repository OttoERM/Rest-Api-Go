package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func SetSubRouters(route *mux.Router) {
	var bookrouter = route.PathPrefix("/books").Subrouter()
	bookrouter.HandleFunc("/new", createBook).Methods("POST")
	bookrouter.HandleFunc("/read/{title}", readBook).Methods("GET")
	bookrouter.HandleFunc("/update", updateBook).Methods("PUT")
	bookrouter.HandleFunc("/delete", deleteBook).Methods("DELETE")
}

func ErrorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
