package main

import "fmt"

func main() {
    numbers := make(map[string]int)
    numbers["one"] = 1
    numbers["ten"] = 10
    numbers["three"] = 3

    fmt.Printf("The third number is: %v\n", numbers["three"]) // get values
    fmt.Printf("numbers: %v\n", numbers)

    delete(numbers, "ten")

    keys := make([]string, len(numbers))
    i := 0
    for k := range numbers {
        keys[i] = k
        i++
    }
    fmt.Printf("keys = %s\n", keys)


    ratings := map[string]float32 { "C":5, "Go":4.5, "Python":4.5, "C++":2 }
    fmt.Printf("ratings: %v\n", ratings)

    csharpRating, ok := ratings["C#"]
    if ok {
        fmt.Println("C# is in the map and its rating is ", csharpRating)
    } else {
        fmt.Println("We have no rating associated with C# in the map")
    }

    r2 := ratings
    r2["C"] = 7
    fmt.Printf("ratings: %v\n", ratings)
}
