package main

import (
	// "bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "os"
)

func main() {
	// ______________________________________________
	// reader := bufio.NewReader(os.Stdin)
	// for {
	// 	fmt.Println("What is your name?")
	// 	text, _ := reader.ReadString('\n')
	// 	fmt.Println("Hello", text)
	// }
	//___________________________________________

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/1", mainPage1)

	port := ":8080"
	fmt.Println("Server listen on port:", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	user := User{"Vasya", "Ivanov"}
	js, _ := json.Marshal(user)
	w.Write(js)
}
func mainPage1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It is amazing"))
}
