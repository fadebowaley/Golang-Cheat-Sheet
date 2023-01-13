type Animal struct {
    Name string
}

type Dog struct {
    Animal
    Breed string
}

d := Dog{Animal{"Fido"}, "Golden Retriever"}
fmt.Println(d.Name)
fmt.Println(d.Breed)