package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

const (
	port = ":9090"
)

func roll(w http.ResponseWriter, r *http.Request) {
	max, err := strconv.Atoi(r.URL.Path[len("/roll/"):])
	if err != nil {
		fmt.Println("somehting bad happened")
	}
	num := rand.Intn(max)
	fmt.Fprintf(w, "%d", num)
}

func main() {
	http.HandleFunc("/roll/", roll)
	http.ListenAndServe(port, nil)
}
