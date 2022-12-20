/*
	Return the number (count) of vowels in the given string.
	We will consider a, e, i, o, u as vowels for this Kata (but not y).
	The input string will only consist of lower case letters and/or spaces.
*/

// my
package kata

import "strings"

func GetCount(str string) (count int) {
	count = 0
	vowels := "aeiou"

	runes := []rune(str)
	for i := 0; i < len(str); i++ {
		if strings.Contains(vowels, string(runes[i])) {
			count++
		}
	}

	return count
}

// bp
package kata

func GetCount(str string) (count int) {
  for _, c := range str {
    switch c {
    case 'a', 'e', 'i', 'o', 'u':
      count++
    }
  }
  return count
}