package main

import (
	"bytes"
	"database/sql"
	"encoding/gob"
	"fmt"
	"log"
	_ "mysql"
	"net"
)

type utenti struct {
	ID      int
	Nome    string
	Cognome string
}

func main() {
	fmt.Println("server listening on 8080")
	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	// listening for incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("new connection")

		// listen to connections in another gorutine
		go listenConnection(conn)
		send(conn)
	}
}

// listening for messages from connection
func listenConnection(conn net.Conn) {
	for {
		buffer := make([]byte, 1400)
		dataSize, err := conn.Read(buffer)
		fmt.Println(dataSize)
		if err != nil {
			fmt.Println("connection closed")
			return
		}

		// the actual message
		data := buffer[:dataSize]
		fmt.Println("received message: ", string(data))

		// echoing the message back out
		_, err = conn.Write(data)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Message sent: ", string(data))
	}
}

func send(conn net.Conn) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/mycinema")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	binBuf := new(bytes.Buffer)

	// create a encoder object
	gobobj := gob.NewEncoder(binBuf)

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	read, err := db.Query("SELECT * FROM utenti")
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer read.Close()

	for read.Next() {
		var utenti utenti
		err := read.Scan(&utenti.ID, &utenti.Nome, &utenti.Cognome)

		if err != nil {
			log.Fatal(err)
		}

		// encode buffer and marshal it into a gob object
		gobobj.Encode(utenti)
		conn.Write(binBuf.Bytes())

		fmt.Printf("%v\n", utenti)
	}
}
