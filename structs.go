package main

import (
    "time"
)

type List struct {
    People []Person
}

type Person struct {
    ID string
    Name string
    FikaTimeStamp time.Time
}

type Tokens struct {
    Auth string
    Created time.Time
}

func (list *List) AddPerson(person Person) []Person {
    list.People = append(list.People, person)
    return list.People
}

func (list *List) UpdateFikaStamp(person Person) []Person {
    return nil
}