package main

import(
    "net/http"
    "log"
    "time"
    "fmt"
    "encoding/json"
    "github.com/gorilla/schema"
    "github.com/nu7hatch/gouuid"
    "github.com/boltdb/bolt"
)

func Ping(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Ping!\n"))
}

func GetNextFika (w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Ping!\n"))
}

func CompletedFika(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Ping!\n"))
}

func NewUser(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    
    if err != nil {
        log.Panicf("Failed to initiate ParseForm: %s", err)
    }
    
    person := new(Person)
    decoder := schema.NewDecoder()
    err = decoder.Decode(person, r.PostForm)
    
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    id, err := uuid.NewV4()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    
    db := OpenDb()    
    list := new(List)
    
    err = db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("Fika"))
        v := b.Get([]byte("List"))
        
        err := json.Unmarshal(v, &list)
        
        if err != nil {
            // Do something
        }
        
        return nil
    })
    
    person.ID = id.String()
    person.FikaTimeStamp = time.Time{}
    list.AddPerson(person)
    
    js, err := json.Marshal(person)
    
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = db.Update(func(tx *bolt.Tx) error {
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
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    CloseDb(db)
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}   

func RemoveUser(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Ping!\n"))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
    db := OpenDb()
    list := new(List)
    
    err := db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("Fika"))
        v := b.Get([]byte("List"))
        
        err := json.Unmarshal(v, &list)
        
        if err != nil {
            // Do something
        }
        
        return nil
    })
    
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    encoded, err := json.Marshal(list)
    
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    CloseDb(db)
    w.Header().Set("Content-Type", "application/json")
    w.Write(encoded)
}