package main

import (
	"fmt"
	"sort"
)

func fillDictionary(words []string) map[string]int {
	dictionary := make(map[string]int)
	for _, word := range words {
		_, ok := dictionary[word]
		if ok {
			dictionary[word]++
		} else {
			dictionary[word] = 1
		}
	}
	return dictionary
}

func topKFrequent(words []string, k int) []string {
	dictionary := fillDictionary(words)
	keys := make([]string, 0, len(dictionary))
	for key := range dictionary {
		keys = append(keys, key)
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)
	sort.SliceStable(keys, func(i, j int) bool {
		return dictionary[keys[i]] > dictionary[keys[j]]
	})
	fmt.Println(keys)

	answer := make([]string, k)

	for i := 0; i < k; i++ {
		answer[i] = keys[i]
	}

	return answer
}

func main() {
	words := []string{"i", "love", "leetcode", "i", "love", "coding"}
	fmt.Println(fillDictionary(words))
	fmt.Println(topKFrequent(words, 2))
	words = []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	fmt.Println(fillDictionary(words))
	fmt.Println(topKFrequent(words, 4))
}
