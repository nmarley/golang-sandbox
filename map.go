package main

import "fmt"

func main() {
    numbers := make(map[string]int)
    numbers["one"] = 1
    numbers["ten"] = 10
    numbers["three"] = 3

    fmt.Printf("The third number is: %v\n", numbers["three"]) // get values
    fmt.Printf("numbers: %v\n", numbers)

    keys := make([]string, len(numbers))
    i := 0
    for k := range numbers {
        keys[i] = k
        i++
    }
    fmt.Printf("keys = %s\n", keys)
}
