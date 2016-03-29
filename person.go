package main

import (
    "fmt"
)

type Person struct {
    Name string
    JobTitle string
    ShoeSize float64
}

type Doctor struct {
    Person
}

type Human interface {
    FullName() string
    SayHello() string
}

type Humans []Human

func (person Person) SayHello() string {
    hello := "Hello, I'm " + person.Name
    return hello
}

func (person Person) GetJobTitle() string {
    fmt.Println("I'm working as a ", person.JobTitle)
    return person.JobTitle
}

func (person Person) FullName() string {
    return person.Name
}

func (doctor Doctor) SayHello() string {
    hello := "Hello, I'm Dr " + doctor.Name
    return hello
}

func (doctor Doctor) FullName() string {
    return doctor.GetJobTitle() + " " + doctor.Name
}

func IteratePeople(people Humans) {
    for _,person := range people {
        if len(person.FullName()) != 0 {
            fmt.Println("f:",person.FullName(), " says '", person.SayHello(), "'.")
        }
    }
}

func main() {
    doctor := new(Doctor)
    doctor.Name = "Meyer"
    doctor.JobTitle = "Dentist"
    doctor.SayHello()
    doctor.GetJobTitle()
    
    walter := new(Person)
    
    people := Humans{}
    people = append(people, *doctor)
    people = append(people, *walter)
    
    IteratePeople(people)
    
    arsling := people[1].(Person)
    arsling.Name = "Arsling"
    people[1] = arsling
    
    IteratePeople(people)
}