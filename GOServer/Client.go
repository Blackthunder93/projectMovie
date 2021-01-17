package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	message = "Test"
)

// QUESTA E UNA MACCHINA DI TEST
func SocketClient(ip string, port int) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer conn.Close()
	conn.Write([]byte(message))
	log.Printf("Send: %s", message)
	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	log.Print(n)
	log.Printf("Receive: %s", buff[:n])
}

func main() {
	var (
		ip   = "127.0.0.1"
		port = 8080
	)
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	defer conn.Close()
	conn.Write([]byte(message))
	log.Printf("Send: %s", message)
	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	log.Print(n)
	log.Printf("Receive: %s", buff[:n])
}