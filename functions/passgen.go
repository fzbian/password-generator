package passGenerator

import "math/rand"

var (
	difficultyEasy = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	difficultyHard = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789!@#$^&*()[]'';{}"
)

func PassGen(lenght int, difficulty int) string {
	var letters []rune
	switch difficulty {
	case 1:
		letters = []rune(difficultyEasy)
	case 2:
		letters = []rune(difficultyHard)
	default:
		letters = []rune(difficultyEasy)
	}
	b := make([]rune, lenght)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
