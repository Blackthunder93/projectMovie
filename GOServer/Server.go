package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Movie struct {
	ID     int    `json:"ID"`
	Titolo string `json:"Titolo"`
	Durata string `json:"Durata"`
}

const (
	connHost       = "localhost"
	connPort       = "8080"
	connType       = "tcp"
	driverName     = "mysql"
	dataSourceName = "root@tcp(127.0.0.1:3306)/test"
)

func main() {
	fmt.Println("Start MAIN")
	SocketServer()
}

// SocketServer
func SocketServer() {
	fmt.Println("Starting " + connType + " server on " + connHost + ":" + connPort)

	listen, err := net.Listen(connType, connHost+":"+connPort)
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
	r := bufio.NewReader(conn)
	n, _ := r.Read(buf)
	data := string(buf[:n])
	log.Printf("Read: %s", data)

	// Write data check
	w := bufio.NewWriter(conn)
	w.Write(getMovie())
	w.Flush()
	log.Printf("Write: %s", getMovie())
}

// Get All Movies from database to json
func getMovie() []byte {
	// Connect to database and run query
	db := connectionDatabase()
	// Read data from database
	read, err := db.Query("select * from film")
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// Create array var
	movieArray := []Movie{}
	// Get multi row
	for read.Next() {
		var movie Movie
		err := read.Scan(&movie.ID, &movie.Titolo, &movie.Durata)
		// If is present error
		if err != nil {
			panic(err.Error())
		}
		// Update array
		movieArray = append(movieArray, movie)
	}

	// Close db and read
	read.Close()
	db.Close()

	// Transform array to json
	respJSON, err := json.Marshal(movieArray)
	if err != nil {
		panic(err.Error())
	}
	// Return json []byte
	return respJSON
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
