// Use a for loop to iterate over an array
arr := []int{1, 2, 3, 4, 5}
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}

// Use a for loop to iterate over a range
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}

// Use a for-each loop to iterate over a slice or map
for key, value := range myMap {
    fmt.Println("Key:", key, "Value:", value)
}