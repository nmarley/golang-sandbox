package main

import "fmt"

type Message struct {
    magic uint32
    command [12]byte
    length uint32
    checksum uint32
    payload []byte
}

func main() {
    // fmt.Println(Vertex{1, 2})
    fmt.Println("hi")
}
