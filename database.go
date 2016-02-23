package main

import (
    "encoding/json"
    "github.com/boltdb/bolt"
)

// Consider moving this to a separate file
func FetchList() (*List, error) {
    db := OpenDb()
    list := new(List) 
    err := db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("Fika"))
        err := b.ForEach(func(k, v []byte) error {
            var fetchedPerson Person
            err := json.Unmarshal(v, &fetchedPerson)
            
            if err != nil {
                return err
            }
            list.AddPerson(&fetchedPerson)
            return nil
        })
        
        if err != nil {
            return err
        }
        
        return nil
    })
    
    if err != nil {
        return nil, err
    }
    CloseDb(db)
    return list, nil
}

func GetOne(name string) (*Person, error) {
    db := OpenDb()
    var p *Person
    err := db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("Fika"))
        v := b.Get([]byte(name))
        
        err := json.Unmarshal(v, &p)
        
        if err != nil {
            return err
        }
        
        return nil
    })
    
    if err != nil {
        return nil, err
    }
    CloseDb(db)
    return p, err
}

func RemoveOne(name string) error {
    db := OpenDb()
    err := db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("Fika"))
        b.Delete([]byte(name))
        
        return nil
    })
    
    if err != nil {
        return err
    }
    CloseDb(db)
    return nil
}

