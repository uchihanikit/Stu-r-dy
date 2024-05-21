package main

import "fmt"

func lengthOfLastWord(s string) int {
	count := 0
	newWord := false
	for _, c := range s {
		if c == ' ' {
			newWord = true
			continue
		} else if newWord {
			count = 0
			newWord = false
		}
		count++
	}
	return count
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
	fmt.Println(lengthOfLastWord("eee e e"))
}
