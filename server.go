package main

import (
    "net/http"
    "fmt"
    "log"
    "time"
    "github.com/gorilla/mux"
    "github.com/boltdb/bolt"
)

func setUpBuckets(db *bolt.DB){
    db.Update(func(tx *bolt.Tx) error {
        _, err := tx.CreateBucketIfNotExists([]byte("Fika"))
        if err != nil {
            return fmt.Errorf("Create bucket: %s", err)
        }
        return nil
    })
}

func OpenDb() *bolt.DB {
    db, err := bolt.Open("my.db", 0644, &bolt.Options{Timeout: 1 * time.Second})
    if err != nil {
        log.Fatal(err)
    }
    
    return db
}

func CloseDb(db *bolt.DB) {
    db.Close()
}

func main() {
    db := OpenDb()
    setUpBuckets(db)
    CloseDb(db)    
    
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/ping", GetOneUser)
    
    r.Path("/fika").
        Methods("GET").
        HandlerFunc(GetNextFika)
        
    r.Path("/fika/{name}").
        Methods("PUT").
        HandlerFunc(CompletedFika)
        
    r.Path("/user").
        Methods("POST").
        HandlerFunc(NewUser)
        
    r.Path("/user").
        Methods("GET").
        HandlerFunc(GetAllUsers)
    
    r.Path("/user/{name}").
        Methods("GET").
        HandlerFunc(GetOneUser)
    
    r.Path("/user/{name}").
        Methods("DELETE").
        HandlerFunc(RemoveUser)
        
    fmt.Println("Server is up and running on localhost:8000")    
    // Bind to a port and pass our router in
    // Nothing will run past this
    http.ListenAndServe(":8000", r)
}