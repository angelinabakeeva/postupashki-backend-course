package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error listening:", err)
        os.Exit(1)
    }
    defer ln.Close()

    for {
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    conn.Write([]byte("OK\n"))
}
