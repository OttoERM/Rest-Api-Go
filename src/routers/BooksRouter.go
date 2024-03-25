package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"RestApiGo/src/db"
	"RestApiGo/src/models"
	"RestApiGo/src/utils"
)

func createBook(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		ErrorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	var e struct {
		Title       string
		Description string
	}

	var DecodedResult = utils.DecodedJSON(r.Body, &e)
	if DecodedResult.HttpCode != 0 {
		ErrorResponse(w, DecodedResult.Message, DecodedResult.HttpCode)
	}

	// var unmarshalErr *json.UnmarshalTypeError

	// decoder := json.NewDecoder(r.Body)
	// decoder.DisallowUnknownFields()
	// err := decoder.Decode(&e)
	// if err != nil {
	// 	if errors.As(err, &unmarshalErr) {
	// 		ErrorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
	// 	} else {
	// 		ErrorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
	// 	}
	// 	return
	// }

	var book = models.Book{
		ID:          -1,
		Title:       e.Title,
		Description: e.Description,
	}
	var statusOperation = db.InsertIntoBook(&book)

	response, err := json.Marshal(&statusOperation)
	if err != nil {
		fmt.Println("Couldn't marshal response")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func readBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Read book"))
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("Update book"))

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
	w.Write([]byte("Delete book"))

}
