package routers

import (
	"net/http"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(203)
	w.Write([]byte("Created User"))
}

func readUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Read User"))
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("Delete User"))

}
