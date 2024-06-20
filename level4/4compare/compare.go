package main

func Compare(word1, word2 string) int {
	for i := 0; i < len(word1) && i < len(word2); i++ {
		if word1[i] < word2[i] {
			return -1
		} else if word1[i] > word2[i] {
			return 1
		}
		if len(word1) < len(word2) {
			return -1
		} else if len(word1) > len(word2) {
			return 1
		}
	}
	return 0
}
