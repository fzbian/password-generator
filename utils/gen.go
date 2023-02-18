package utils

import (
	"math/rand"
	"time"
)

var (
	characters = "abcdefghijklmnñopqrfP8ruvDkjcqemnWXatLREqiy2heFeMBTcZ123456789!@#$^&*"
)

func Gen(lenght int) string {
	rand.Seed(time.Now().UnixNano())
	var letters []rune
	letters = []rune(characters)
	b := make([]rune, lenght)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
