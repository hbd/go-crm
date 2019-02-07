package handlers

import (
     "encoding/json"
     "net/http"
     "fmt"
     "github.com/gorilla/mux"   
     
     "database/sql"
     _ "github.com/lib/pq"

     )

type Person struct {
     ID           string    `json:"id,omitempty"`
     Firstname    string    `json:"firstname,omitempty"`
     Lastname     string    `json:"lastname,omitempty"`
     Status       string    `json:"status,omitempty`
     Address      *Address  `json:"address,omitempty"`
}

type Address struct {
   City           string    `json:"city,omitempty"`
   State          string    `json:"state,omitempty"`
}

var people []Person

func GetClients(w http.ResponseWriter, r *http.Request){
     db, _ = sql.OpenDB("postgres")
     query_res := db.Exec(`
     select * from client`)
     fmt.Printf("%s\n", query_res)
     json.NewEncoder(w).Encode(people)
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
     json.NewEncoder(w).Encode(people)
}
func GetPerson(w http.ResponseWriter, r *http.Request) {
     params := mux.Vars(r)
     for _, item := range people {
          if item.ID == params["id"] {
               json.NewEncoder(w).Encode(item)
               return
          }
     }
     json.NewEncoder(w).Encode(&Person{})
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
     params := mux.Vars(r)
     for _, item := range people {
          if item.ID == params["id"] {
               json.NewEncoder(w).Encode(item.Status)
               return
          }
     }
     json.NewEncoder(w).Encode("")
}

func SetStatus(w http.ResponseWriter, r *http.Request) {
     params := mux.Vars(r)
     for _, item := range people {
          if item.ID == params["id"] {
               item.Status = params["newstatus"]
               json.NewEncoder(w).Encode(people)
               return
          }
     }
     json.NewEncoder(w).Encode("User with given id not found")
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
     params := mux.Vars(r)
     var person Person
     _ = json.NewDecoder(r.Body).Decode(&person)
     person.ID = params["id"]
     people = append(people, person)
     json.NewEncoder(w).Encode(people)
}
func DeletePerson(w http.ResponseWriter, r *http.Request) {
}