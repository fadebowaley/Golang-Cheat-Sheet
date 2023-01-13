// Declare an array with a fixed size
var a [5]int
a[0] = 1
a[1] = 2
fmt.Println(a)

// Declare an array with a composite literal
b := [3]int{1, 2, 3}
fmt.Println(b)

// Iterate over an array
for i, v := range a {
    fmt.Println(i, v)
}