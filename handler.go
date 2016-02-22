package main

import(
    "net/http"
    "log"
    "time"
    "encoding/json"
    "github.com/gorilla/schema"
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
    
    person.FikaTimeStamp = time.Time{}
    js, err := json.Marshal(person)
    
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Ping!\n"))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Ping!\n"))
}