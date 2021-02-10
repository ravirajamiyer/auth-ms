package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "database/sql"
    _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "****"
  dbname   = "authy"
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
    var  (
       id string
       email string
    )

    var users [] User

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

    fmt.Println("Openining DB connection")
    db,err := sql.Open("postgres", psqlInfo)
    if err != nil {
      panic(err)
    }
    defer db.Close()
    fmt.Println("Opened DB connection")

    fmt.Println("Successfully connected!")

    rows, err := db.Query("select id, email from users")
    for rows.Next() {
	err := rows.Scan(&id, &email)
	if err != nil {
		log.Fatal(err)
	}
        var user = User{id: id, email: email}
        users = append(users,user)
    }
	//fmt.Println(users)
        json.NewEncoder(w).Encode(users)
}

func returnSingleUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    for _, user := range users {
        if user.id == key {
            json.NewEncoder(w).Encode(user)
        }
    }
}

func main() {

    //users = []User{
    //    User{id: "1", email: "user1@gmail.com"},
    //    User{id: "2", email: "user2@gmail.com"},
    //}

    handleRequests()
}

type Message struct {
    Message string `json:"Message"`
}



type User struct {
    id      string `json:"id"`
    email string `json:"email"`
}

var users []User
var message Message
