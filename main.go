package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	var (
		minInt int
		maxInt int
	)

	fmt.Println("Type Min integer")
	fmt.Fscan(os.Stdin, &minInt)
	fmt.Println("Type Max integer")
	fmt.Fscan(os.Stdin, &maxInt)

	rand.Seed(time.Now().UTC().UnixNano())
	randomNumber := strconv.Itoa(minInt + rand.Intn(maxInt-minInt+1))
	fmt.Println("Your random number is:", randomNumber)

	http.HandleFunc("/random", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(randomNumber))
	})

	port := ":8080"
	fmt.Println("Server listen on port:", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
