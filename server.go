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

func main() {
    db, err := bolt.Open("my.db", 0644, &bolt.Options{Timeout: 1 * time.Second})
    if err != nil {
        log.Fatal(err)
    }
    
    defer db.Close()
    setUpBuckets(db)
    
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/ping", Ping)
    
    r.HandleFunc("/fika", GetNextFika).
        Methods("GET")
        
    r.HandleFunc("/fika", CompletedFika).
        Methods("PUT")
        
    r.HandleFunc("/user", NewUser).
        Methods("POST")
        
    r.HandleFunc("/user", RemoveUser).
        Methods("DELETE")
        
    r.HandleFunc("/user", GetAllUsers).
        Methods("GET")
    

    fmt.Println("Server is up and running on localhost:8000")    
    // Bind to a port and pass our router in
    // Nothing will run past this
    http.ListenAndServe(":8000", r)
}