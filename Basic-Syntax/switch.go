// Use a switch statement
switch x := 5; x {
case 1:
    fmt.Println("x is 1")
case 2:
    fmt.Println("x is 2")
default:
    fmt.Println("x is not 1 or 2")
}

// Use a switch statement with multiple cases
switch x := 5; x {
case 1, 2:
    fmt.Println("x is 1 or 2")
default:
    fmt.Println("x is not 1 or 2")
}

// Use a switch statement with expressions
switch {
case x > 0:
    fmt.Println("x is positive")
case x < 0:
    fmt.Println("x is negative")
default:
    fmt.Println("x is zero")
}