// Declare a map with a composite literal
m := map[string]int{"key1": 1, "key2": 2}
fmt.Println(m)

// Declare a map with the make function
n := make(map[string]int)
n["key1"] = 1
n["key2"] = 2
fmt.Println(n)

// Get the value of a key
fmt.Println(m["key1"])

// Check if a key exists in a map
value, exists := m["key3"]
if exists {
    fmt.Println(value)
} else {
    fmt.Println("key does not exist")
}