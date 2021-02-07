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
    router.HandleFunc("/user/{id}", returnSingleUser)
    log.Fatal(http.ListenAndServe(":8080", router))
}

func returnAllUsers(w http.ResponseWriter, r *http.Request){
    fmt.Println("Returning all users")
    json.NewEncoder(w).Encode(users)
}

func returnSingleUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    // Loop over all of our Articles
    // if the article.Id equals the key we pass in
    // return the article encoded as JSON
    for _, user := range users {
        if user.Id == key {
            json.NewEncoder(w).Encode(user)
        }
    }
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
