// Define a struct
type Person struct {
    Name string
    Age  int
}

// Create an instance of the struct
p := Person{Name: "John", Age: 30}

// Access struct fields
fmt.Println(p.Name)
fmt.Println(p.Age)

// Modify struct fields
p.Name = "Jane"
p.Age = 25

// Declare a struct with embedded fields
type Employee struct {
    Person
    Role string
}

// Create an instance of the struct
e := Employee{Person: Person{Name: "John", Age: 30}, Role: "Manager"}

// Access embedded fields
fmt.Println(e.Name)
fmt.Println(e.Age)
fmt.Println(e.Role)