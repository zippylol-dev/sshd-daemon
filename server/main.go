package main

import (
	"log"
	"net"
	"os"
	"os/exec"
	"io"
	"github.com/creack/pty"
)

func main() {
	log.Println("Starting mysshd on :2222")

	ln, err := net.Listen("tcp", ":2222")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Println("Listening...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Accept error:", err)
			continue
		}
		log.Printf("Accepted connection from %s", conn.RemoteAddr())
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	shell := os.ExpandEnv("$HOME/adb/rish")
	cmd := exec.Command(shell)

	// Create a PTY
	ptmx, err := pty.Start(cmd)
	if err != nil {
		log.Printf("PTY error: %v", err)
		conn.Write([]byte("Failed to start shell.\n"))
		return
	}
	defer ptmx.Close()

	// Bidirectional copy between PTY and network connection
	go func() {
		_, _ = io.Copy(ptmx, conn) // stdin
	}()
	_, _ = io.Copy(conn, ptmx) // stdout/stderr
}
