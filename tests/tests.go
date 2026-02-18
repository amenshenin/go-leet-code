package tests

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/amenshenin/go-leet-code/tests/easy"
)

type Test interface {
	Run() error
}

var easytests = map[string]Test{
	"first": &easy.First{},
}
var mediumtests = map[string]Test{}
var hardtests = map[string]Test{}

func Run(grade string, name string) error {
	source, error := getGrade(grade)
	if error != nil {
		return error
	}
	test, ok := source[name]
	if !ok {
		return fmt.Errorf("Unsupported test: %s", name)
	}
	test.Run()
	return nil
}

func GetNewName(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func getGrade(grade string) (map[string]Test, error) {
	source := map[string]Test{}
	switch grade {
	case "easy":
		source = easytests
	case "medium":
		source = mediumtests
	case "hard":
		source = hardtests
	default:
		return source, fmt.Errorf("Unavailable grade: %s", grade)
	}
	return source, nil
}
