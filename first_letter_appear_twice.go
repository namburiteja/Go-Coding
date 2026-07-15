package main

import "fmt"

func repeatedCharacter(s string) byte {
	length := len(s)
	hashMap := make(map[byte]struct{})

	for i := 0; i < length; i++ {
		char := s[i]

		if _, exists := hashMap[char]; exists {
			return char
		}

		hashMap[char] = struct{}{}
	}

	return 'f'
}

func main() {
	s := "abccbaacz"

	result := repeatedCharacter(s)

	if result == 'f' {
		fmt.Println("No repeated character found.")
	} else {
		fmt.Printf("First repeated character: %c\n", result)
	}
}