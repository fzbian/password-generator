package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

const (
	port = ":3000"
)

var (
	difficultyEasy = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	difficultyHard = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789!@#$^&*()[]'';{}"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api", processor)
	fmt.Println("Server started on port " + port)
	http.ListenAndServe(port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ok.")
}

type Body struct {
	Lenght     int
	Difficulty int
}

func processor(w http.ResponseWriter, r *http.Request) {
	body := r.URL.Query()
	x, _ := strconv.Atoi(body.Get("lenght"))
	a, _ := strconv.Atoi(body.Get("difficulty"))
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["password"] = passGen(x, a)
	jsonResponse, _ := json.Marshal(resp)
	w.Write(jsonResponse)
	return
}

func passGen(lenght int, difficulty int) string {
	switch difficulty {
	case 1:
		var letters = []rune(difficultyEasy)
		b := make([]rune, lenght)

		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		return string(b)
	case 2:
		var letters = []rune(difficultyHard)
		b := make([]rune, lenght)

		for i := range b {
			b[i] = letters[rand.Intn(len(letters))]
		}
		return string(b)
	default:
		return string("This difficulty does not exist")
	}
}
