package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"RestApiGo/src/models"
)

var DbConn, err = sql.Open("mysql", "root:<password>@tcp(127.0.0.1:3306)/mydatabaseforgo?parseTime=true")

func ConnectDb() {
	if err != nil {
		panic(err.Error())
	}

	err = DbConn.Ping()

	if err != nil {
		panic("Couldn't connect to mysql")
	}
}

func CreateTableUser() {
	var query = `CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`
	_, err = DbConn.Exec(query)
}

type statusOperation struct {
	Status  string
	IsError bool
}

func InserIntoUser(user *models.User) {

}

func InsertIntoBook(book *models.Book) *statusOperation {
	var query = `INSERT INTO books
	(Title, Description) VALUES (?, ?)`

	_, err = DbConn.Exec(query, book.Title, book.Description)

	if err != nil {
		return &statusOperation{Status: err.Error(), IsError: true}
	}

	return &statusOperation{Status: "Book created succesfully", IsError: false}
}
