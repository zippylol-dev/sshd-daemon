package main

import (
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	serverAddr := os.Getenv("MYSSH_SERVER")
	if serverAddr == "" {
		serverAddr = "127.0.0.1:2222"
		log.Println("MYSSH_SERVER not set, defaulting to 127.0.0.1:2222")
	}

	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		conn.Close()
		os.Exit(0)
	}()

	// Pipe input/output
	go io.Copy(conn, os.Stdin)
	io.Copy(os.Stdout, conn)
}
