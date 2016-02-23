package main

import (
    "time"
    "fmt"
    "strings"
)

const timestampFormat = "02 Jan 2006 15:04"

func Equals(a string, b string) bool {
    return strings.ToLower(a) == strings.ToLower(b) 
}

// My own time wrapper
type Timestamp time.Time
func (t Timestamp) MarshalJSON() ([]byte, error) {
    fmt.Printf("Marshalling to JSON: \n %s", t.ToString())
    stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(timestampFormat))
    return []byte(stamp), nil
}

func (t *Timestamp) UnmarshalJSON(raw []byte) error {
    stamp, _ := time.Parse(timestampFormat, string(raw))
    fmt.Printf("Unmarshalling from JSON: \n %s ", stamp.String())
    *t = Timestamp(stamp)
    return nil
}

func (t Timestamp) ToString() (string) {
    return fmt.Sprintf("\"%s\"", time.Time(t).Format(timestampFormat))
}

func (t Timestamp) Parse() (time.Time, error) {
    timestamp, err := time.Parse(timestampFormat, t.ToString())
    
    if err != nil {
        return time.Time{}, err
        // Do something
    }
    
    return timestamp, nil
}