// Define a function that returns an error
func divide(x, y int) (int, error) {
    if y == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return x / y, nil
}

// Use the function and check for errors
result, err := divide(5, 2)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(result)
}