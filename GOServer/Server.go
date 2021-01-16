package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

type Users struct {
	Id      int `json:"Id"`
	FirstName    string `json:"Firstname"`
	LastName string `json:"Lastname"`
}

// Connection database and get rows
func connectionDatabase(query string) (*sql.Rows) {
	// Connection to database
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/mycinema")
	defer db.Close()
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	// Read data from database
	read, err := db.Query(query)
	defer read.Close()
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	return read
}

func main() {
	// Go to Api
	http.HandleFunc("/api/v1/users", getUsers)
	// If is present error
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// Get All Users from database
func getUsers(w http.ResponseWriter, r *http.Request) {
	// Connect to database and run query
	read := connectionDatabase("SELECT * FROM utenti")
	// Create array var
	usersArray := []Users{}
	// Get multi row
	for read.Next() {
		var users Users
		err := read.Scan(&users.Id, &users.FirstName, &users.LastName)
		// If is present error
		if err != nil {
			// Get error 500
			w.WriteHeader(500)
		}
		// Update array
		usersArray = append(usersArray, users)
	}
	// Create Json
	respJson, err := json.Marshal(usersArray)
	if err != nil {
		// Get error 500
		w.WriteHeader(500)
	}
	// Write json and stamp
	r.Header.Add("content-type","application/json")
	w.Write(respJson)
}
