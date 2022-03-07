package main

import (
	"encoding/gob"
	"faraway/pkg/blockchain"
	"fmt"
	"io/ioutil"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "server:8081")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	bc := blockchain.NewBlockchain()
	bc.AddBlock("Block 1")
	bc.AddBlock("Block 2")

	enc := gob.NewEncoder(conn)
	err = enc.Encode(bc)
	if err != nil {
		panic(err)
	}

	resp, err := ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resp))
}
