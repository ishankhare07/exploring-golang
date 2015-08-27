package main

import "fmt"

type Person struct {
    Name string
}

func (p *Person) talk() {
    fmt.Println("hello my name is : ", p.Name)
}

type Android struct {
    Person
    model string
}

func main() {
    a := Android{
        Person{"ishan khare"},
        "some model"}
    //a.Person.Name = "ishan khare"
    a.talk()
}
