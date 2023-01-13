// Declare a variable
x := 5

// Declare a pointer to the variable
var p *int = &x

// Dereference the pointer to get the value
fmt.Println(*p)

// Modify the value through the pointer
*p = 10
fmt.Println(x)