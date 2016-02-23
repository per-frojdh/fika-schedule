package main

import(
    "net/http"
    "log"
    "fmt"
    "time"
    "github.com/gorilla/mux"
    "github.com/gorilla/schema"   
)

func Ping(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Ping!\n"))
}

func GetNextFika (w http.ResponseWriter, r *http.Request) {
    list := new(List)
    list, err := FetchList()
    
    if err != nil {
        http.Error(w, "Failed to fetch list", http.StatusInternalServerError)
        return
    }
    
    var fika Person
    highest := time.Time{}
    
    for i := len(list.People) - 1; i >= 0; i-- {
        current := list.People[i]
        fmt.Printf("Comparing (%s) %s (current) with %s (highest) \n", current.Name, current.FikaTimeStamp.String(), highest.String())
        if err != nil {
            // Do something
        }
        
        if !current.FikaTimeStamp.After(highest) {
            fika = current
        }
    }
    
    json, err := fika.ToJSON()
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(json)
}

func CompletedFika(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]
    if len(name) <= 0 {
        http.Error(w, "Missing name in url", http.StatusBadRequest)
        return
    }
    
    // Should probably check that we have one first
    // Also we create a new empty object to database sometimes. "Fun"
    
    target, err := GetOne(name)
    target.FikaTimeStamp = time.Now()
    target.Save()

    if err != nil {
        http.Error(w, "Missing name in url", http.StatusBadRequest)
        return
    }
    
    js, err := target.ToJSON()
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
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
    
    if err != nil {
        http.Error(w, "Failed to fetch list", http.StatusInternalServerError)
        return
    }

    err = person.Initialize()
    
    if err != nil {
        http.Error(w, "Failed to initiate person", http.StatusInternalServerError)
        return
    }
    
    // list.AddPerson(person)
    // list.Save()
    person.Save()
    
    js, err := person.ToJSON()
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}   

func RemoveUser(w http.ResponseWriter, r *http.Request) {    
    vars := mux.Vars(r)
    name := vars["name"]
    if len(name) <= 0 {
        http.Error(w, "Missing name in url", http.StatusBadRequest)
        return
    }
    
    RemoveOne(name)    
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte("Successfully deleted " + name))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
    list := new(List)
    list, err := FetchList()
    
    if err != nil {
        http.Error(w, "Failed to fetch list", http.StatusInternalServerError)
    }
    
    json := list.ToJSON()
    
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(json)
}

func GetOneUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]
    if len(name) <= 0 {
        http.Error(w, "Missing name in url", http.StatusBadRequest)
        return
    }
    
    p, err := GetOne(name)
    
    if err != nil {
        http.Error(w, "Failed to fetch list", http.StatusInternalServerError)
    }
    
    json, err := p.ToJSON()
    
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(json)
}