// Define an interface
type Shape interface {
    Area() float64
}

// Define a struct that implements the interface
type Rectangle struct {
    width  float64
    height float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}

// Use the interface
var s Shape = Rectangle{width: 10, height: 5}
fmt.Println(s.Area())


_, err := os.Open("file.txt")
if err != nil {
    fmt.Println(err)
}