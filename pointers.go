package main

import "fmt"

func main() {
    var p *int
    var i int = 42
    p = &i

    *p = 21

    fmt.Println(*p)
}
