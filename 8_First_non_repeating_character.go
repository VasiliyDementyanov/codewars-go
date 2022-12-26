/*
Write a function named first_non_repeating_letter that takes a string input, 
and returns the first character that is not repeated anywhere in the string.

For example, if given the input 'stress', 
the function should return 't', since the letter t only occurs once in the string, and occurs first in the string.

As an added challenge, upper- and lowercase letters 
are considered the same character, but the function should return the correct case for the initial letter. 
For example, the input 'sTreSS' should return 'T'.

If a string contains all repeating characters, 
it should return an empty string ("") or None -- see sample tests.
*/

// my
package kata
import "strings"

func FirstNonRepeating(str string) string {
  if str == "" {
    return ""
  }
  
  runeSample := []rune(str)
  strLower := strings.ToLower(str)
  
  minLetter := ""
  
  for i := 0; i < len(runeSample); i++ {
    count := strings.Count(strLower, strings.ToLower(string(runeSample[i])))
    
    if count == 1 && minLetter == "" {
      minLetter = string(runeSample[i])
      break
    }
  }
  
	return minLetter
}

// bp
package kata
import (
	"strings"
)
func FirstNonRepeating(str string) string {
    for _, c := range str {
        if strings.Count(strings.ToLower(str), strings.ToLower(string(c))) < 2 {
	          return string(c)
	      }
    }
    return ""
}