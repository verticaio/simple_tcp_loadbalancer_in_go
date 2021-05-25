package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

var (
	counter int

	// TODO configurable
	listenAddr = ":8080"

	// TODO configurable
	server = []string{
		"localhost:5001",
		"localhost:5002",
		"localhost:5003",
	}
)

func main() {

	listener, err := net.Listen("tcp",listenAddr)
	if err != nil {
		log.Fatalf("failed to listen : %s", err)
	}

	defer listener.Close()  // It means codes is working after that line

	// forever loop for accept connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept cnnection : %s", err)
		}

		backend :=  chooseBackend()
		fmt.Printf("counter=%d, backedn=%s\n", counter, backend)
		go func() {
			err := proxy(backend, conn)
			if err != nil {
				log.Printf("WARNING: proxying fail: %v", err)
			}
		}()
	}

}

func proxy(backend string, c net.Conn) error {
	bc , err := net.Dial("tcp", backend)
    if err != nil {
    	return fmt.Errorf("failed to connect backend %s: %v", backend, err)
    }
    // c --> bc   connections to backend
    go io.Copy(bc, c)   // send to background  with goroutine
    // bc --> c   get result to client
    go io.Copy(c,bc)    // send to background  with goroutine

    return nil

}

func chooseBackend() string{
	// TODO choose randomly, round robin
	//fmt.Println(counter , len(server), counter%len(server))
	//fmt.Printf("\n")
	s := server[counter%len(server)]  // len(server) = 3 |  0 , 1 , 2
	counter++
	return  s
}
