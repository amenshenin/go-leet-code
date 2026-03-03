package medium

import (
	"fmt"
	"math/rand"
	"time"
)

type Y7FRc struct {
	IsComplete int
	Name       string
	Descr      string
}

func Y7FRcInit() *Y7FRc {
	return &Y7FRc{
		IsComplete: 0,
		Name:       "https://leetcode.com/problems/longest-substring-with-at-most-k-distinct-characters/description/",
		Descr:      `Дана строка s и целое число k. Верните длину самой длинной подстроки s, которая содержит не более k различных символов.`,
	}
}

func (f *Y7FRc) Run() error {
	fmt.Printf("%+v\n", f)
	var str = getString(50)
	var k = 15
	fmt.Printf("\nСтрока: %s\nЧисло: %d\n", str, k)

	runes := []rune(str)
	strlen := len(runes)
	if getLengthOfLongestSubstringKDistinct(runes) <= k {
		fmt.Printf("Длинна равна %d", strlen)
		return nil
	}
	// j := strlen - 1
	// for i := 0; i < j; i++ {

	// 	for j; j < i; j++ {

	// 	}
	// }

	return nil
}

func getLengthOfLongestSubstringKDistinct(str []rune) int {
	charCount := make(map[rune]int)
	for _, val := range str {
		_, ok := charCount[val]
		if !ok {
			charCount[val] = 1
		}
	}
	return len(charCount)
}

func getString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
