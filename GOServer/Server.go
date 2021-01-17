package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net"
	"os"
)

type Users struct {
	Id      int `json:"Id"`
	FirstName    string `json:"Firstname"`
	LastName string `json:"Lastname"`
}

const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
	driverName = "mysql"
	dataSourceName = "root@tcp(127.0.0.1:3306)/mycinema"
)

func main() {
	fmt.Println("Start MAIN")
	SocketServer()
}

// SocketServer
func SocketServer() {
	fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)

	listen, err := net.Listen(connType, connHost + ":" + connPort)
	if err != nil {
		log.Fatalf("Socket listen port %d failed,%s", connPort, err)
		os.Exit(1)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handler(conn)
	}
}

// handler method
func handler(conn net.Conn) {
	defer conn.Close()

	// Read data check
	buf := make([]byte, 1024)
	r   := bufio.NewReader(conn)
	n, _ := r.Read(buf)
	data := string(buf[:n])
	log.Printf("Read: %s", data)

	// Write data check
	w := bufio.NewWriter(conn)
	w.Write(getUsers())
	w.Flush()
	log.Printf("Write: %s", getUsers())
}

// Get All Users from database to json
func getUsers() []byte {
	// Connect to database and run query
	db := connectionDatabase()
	// Read data from database
	read, err := db.Query("select * from utenti")
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// Create array var
	usersArray := []Users{}
	// Get multi row
	for read.Next() {
		var users Users
		err := read.Scan(&users.Id, &users.FirstName, &users.LastName)
		// If is present error
		if err != nil {
			panic(err.Error())
		}
		// Update array
		usersArray = append(usersArray, users)
	}

	// Close db and read
	read.Close()
	db.Close()

	// Transform array to json
	respJson, err := json.Marshal(usersArray)
	if err != nil {
		panic(err.Error())
	}
	// Return json []byte
	return respJson
}

// Connection database and get rows
func connectionDatabase() *sql.DB {
	// Connection to database
	db, err := sql.Open(driverName, dataSourceName)
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	return db
}
