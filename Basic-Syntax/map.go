package main

import "fmt"

func main() {
    // Creating a map using a composite literal
    a := map[string]int{"one": 1, "two": 2, "three": 3}
    fmt.Println("a:", a)

    // Creating a map using the make function
    b := make(map[string]int)
    b["one"] = 1
    b["two"] = 2
    b["three"] = 3
    fmt.Println("b:", b)

    // Accessing elements in a map
    fmt.Println("a[\"two\"]:", a["two"])

    // Modifying elements in a map
    a["two"] = 20
    fmt.Println("a:", a)

    // Deleting elements in a map
    delete(a, "two")
    fmt.Println("a:", a)

    // Checking if a key is in a map
    value, exists := a["two"]
    if  !exists {
        fmt.Println("a[\"two\"]:", value)
    } else {
        fmt.Println("a[\"two\"] does not exist.")
    }
}