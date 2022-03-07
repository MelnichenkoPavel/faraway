package main

import (
	"encoding/gob"
	"faraway/pkg/blockchain"
	"fmt"
	"log"
	"math/rand"
	"net"
)

var wow = []string{
	"“The best way out is always through.”- Robert Frost",
	"“Always Do What You Are Afraid To Do” – Ralph Waldo Emerson",
	"“Believe and act as if it were impossible to fail.” – Charles Kettering",
}

func main() {

	fmt.Println("Launching server...")

	// Устанавливаем прослушивание порта
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {

	defer conn.Close()

	var bc blockchain.Blockchain
	dec := gob.NewDecoder(conn)
	err := dec.Decode(&bc)
	if err != nil {
		log.Println(err)
		return
	}

	if !bc.Validate() {
		log.Println("Invalid blockchain", bc)
		return
	}

	quote := wow[rand.Uint64()%uint64(len(wow))]

	// Отправить новую строку обратно клиенту
	_, err = conn.Write([]byte(quote))
	if err != nil {
		log.Println(err)
	}
}
