package main

import (
    "time"
    "fmt"
    "errors"
    "encoding/json"
    "github.com/nu7hatch/gouuid"
    "github.com/boltdb/bolt"
)

type List struct {
    People []Person
}

type Person struct {
    ID string
    Name string
    Created time.Time
    FikaTimeStamp time.Time
}

func (person *Person) Initialize() error {
    id, err := uuid.NewV4()
    if err != nil {
        return err
    }
    
    person.ID = id.String()
    person.Created = time.Now()
    person.FikaTimeStamp = time.Time{}
    
    return nil
}

func (person *Person) Save() (*Person, error) {
    db := OpenDb()
    err := db.Update(func(tx *bolt.Tx) error {
        b, err := tx.CreateBucketIfNotExists([]byte("Fika"))
        if err != nil {
            return fmt.Errorf("Create bucket: %s", err)
        }
        encoded, err := json.Marshal(person)
        err = b.Put([]byte(person.Name), encoded)
        if err != nil {
            return fmt.Errorf("Error inserting data: %s", err)
        }
        return nil
    })
    
    if err != nil {
        return nil, err
    }
    
    CloseDb(db)
    return person, nil
}

func (person *Person) ToJSON() ([]byte, error) {
    encoded, err := json.Marshal(person)
    
    if err != nil {
        return nil, err
        // Do something
    }
    
    return encoded, nil
}

type Tokens struct {
    Auth string
    Created time.Time
}

func (list *List) AddPerson(person *Person) []Person {
    list.People = append(list.People, *person)
    return list.People
}

func (list *List) GetPerson(name string) (Person, error) {
    for i := len(list.People) - 1; i >= 0; i-- {
        current := list.People[i]
        if Equals(current.Name, name) {
            return current, nil
        }
    }
    return Person{}, errors.New("Couldn't find person")
}

func (list *List) RemovePerson(name string) bool {
    nameFound := false
    for i := len(list.People) - 1; i >= 0; i-- {
        personToDelete := list.People[i]
        
        if Equals(personToDelete.Name, name) {
            nameFound = true
            list.People = append(list.People[:i], list.People[i+1:]...)
        }
    }
    return nameFound
}

func (list *List) ToJSON() []byte {
    encoded, err := json.Marshal(list)
    
    if err != nil {
        // Do something
    }
    
    return encoded
}

func (list *List) Update(person Person) ([]Person, error) {
    for i := len(list.People) - 1; i >= 0; i-- {
        updateTarget := list.People[i]
        
        if Equals(updateTarget.Name, person.Name) {
            list.People[i] = person
        }
    }
    return list.People, nil
}

func (list *List) Save() ([]Person, error) {
    db := OpenDb()
    err := db.Update(func(tx *bolt.Tx) error {
        b, err := tx.CreateBucketIfNotExists([]byte("Fika"))
        if err != nil {
            return fmt.Errorf("Create bucket: %s", err)
        }
        encoded, err := json.Marshal(list)
        err = b.Put([]byte("List"), encoded)
        
        if err != nil {
            return fmt.Errorf("Error inserting data: %s", err)
        }
        return nil
    })
    
    if err != nil {
        return nil, err
    }
    
    CloseDb(db)
    return list.People, nil
}