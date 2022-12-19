/*
Write a function that takes in a string of one or more words, and returns the same string,
but with all five or more letter words reversed (Just like the name of this Kata).
Strings passed in will consist of only letters and spaces. Spaces will be included only when more than one word is present.

spinWords( "Hey fellow warriors" ) => returns "Hey wollef sroirraw"
spinWords( "This is a test") => returns "This is a test"
spinWords( "This is another test" )=> returns "This is rehtona test"

*/

// my
package kata

import "strings"

func SpinWords(str string) string {
	sArr := strings.Split(str, " ")
	var reversedArr []string

	for _, word := range sArr {
		if len(word) > 4 {
			reversedWord := []rune(word)
			for i, j := 0, len(reversedWord)-1; i < j; i, j = i+1, j-1 {
				reversedWord[i], reversedWord[j] = reversedWord[j], reversedWord[i]
			}
			reversedArr = append(reversedArr, string(reversedWord))
		} else {
			reversedArr = append(reversedArr, word)
		}
	}
	return strings.Join(reversedArr, " ")
}

// bp

func SpinWords(str string) string {
	aTokens := strings.Split(str, " ")
	for i, word := range aTokens {
		if len([]rune(word)) > 4 {
			aTokens[i] = reverse(word)
		}
	}
	return strings.Join(aTokens, " ")
} // SpinWords

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
