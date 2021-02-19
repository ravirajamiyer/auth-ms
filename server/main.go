package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "authy"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
	message = Message{Message: "Hello, welcome"}
	json.NewEncoder(w).Encode(message)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/users", returnAllUsers)
	router.HandleFunc("/user/{id}", returnSingleUser)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	var (
		id    string
		email string
	)

	var users []User

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	fmt.Println("Openining DB connection")
	db, err := sql.Open("postgres", psqlInfo)
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
		var user = User{Id: id, Email: email}
		users = append(users, user)
	}
	//fmt.Println(users)
	json.NewEncoder(w).Encode(users)
}

func returnSingleUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println("Key: " + key)
	var id string
	var email string

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	fmt.Println("Openining DB connection")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	fmt.Println("Opened DB connection")
	rows, err := db.Query("select id,email from users where id=$1", key)
	defer rows.Close()
	for rows.Next() {
		switch err := rows.Scan(&id, &email); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			fmt.Printf("Data row = (%s, %s)\n", id, email)
			var user = User{Id: id, Email: email}
			json.NewEncoder(w).Encode(user)
		default:
			checkError(err)
		}
	}

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	handleRequests()
}

type Message struct {
	Message string `json:"Message"`
}

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

var users []User
var message Message
