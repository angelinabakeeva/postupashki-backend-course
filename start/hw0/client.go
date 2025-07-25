package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting:", err)
        os.Exit(1)
    }
    defer conn.Close()

    response, err := bufio.NewReader(conn).ReadString('\n')
    if err != nil {
        fmt.Println("Error reading response:", err)
        os.Exit(1)
    }

    if strings.TrimSpace(response) != "OK" {
        fmt.Println("Unexpected response:", response)
        os.Exit(1)
    }

    fmt.Println("Server responded correctly:", response)
}
