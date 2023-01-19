package main

import "fmt"

func main() {
    // Creating a slice using the make function
    a := make([]int, 5)
    fmt.Println("a:", a)

    // Creating a slice using a literal
    b := []int{1, 2, 3, 4, 5}
    fmt.Println("b:", b)

    // Creating a slice using a composite literal
    c := []int{1, 2, 3, 4, 5}
    fmt.Println("c:", c)

    // Accessing elements in a slice
    fmt.Println("b[2]:", b[2])

    // Modifying elements in a slice
    b[2] = 10
    fmt.Println("b:", b)

    // Slicing a slice
    d := b[1:3]
    fmt.Println("d:", d)

    // Appending to a slice
    e := append(d, 15)
    fmt.Println("e:", e)

    // Creating a slice with a capacity
    f := make([]int, 0, 5)
    fmt.Println("f:", f)
}