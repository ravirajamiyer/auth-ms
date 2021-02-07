package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: homePage")
    message = Message {Message: "Hello, welcome"}
    json.NewEncoder(w).Encode(message)
}

func handleRequests() {
    router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/", homePage)
    router.HandleFunc("/users", returnAllUsers)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func returnAllUsers(w http.ResponseWriter, r *http.Request){
    json.NewEncoder(w).Encode(users)
}

func main() {
    users = []User{
        User{Id: "1", FirstName: "User1", LastName: "FirstUser", Email: "user1@gmail.com"},
        User{Id: "2", FirstName: "User2", LastName: "2nduser", Email: "user2@gmail.com"},
    }

    handleRequests()
}

type Message struct {
    Message string `json:"Message"`
}



type User struct {
    Id      string `json:"Id"`
    FirstName string `json:"FirstName"`
    LastName string `json:"LastName"`
    Email string `json:"Email"`
}

var users []User
var message Message
