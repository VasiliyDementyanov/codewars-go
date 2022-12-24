/*
You are given an array(list) strarr of strings and an integer k. 
Your task is to return the first longest string consisting of k consecutive strings taken in the array.

Examples:
strarr = ["tree", "foling", "trashy", "blue", "abcdef", "uvwxyz"], k = 2

Concatenate the consecutive strings of strarr by 2, we get:

treefoling   (length 10)  concatenation of strarr[0] and strarr[1]
folingtrashy ("      12)  concatenation of strarr[1] and strarr[2]
trashyblue   ("      10)  concatenation of strarr[2] and strarr[3]
blueabcdef   ("      10)  concatenation of strarr[3] and strarr[4]
abcdefuvwxyz ("      12)  concatenation of strarr[4] and strarr[5]

Two strings are the longest: "folingtrashy" and "abcdefuvwxyz".
The first that came is "folingtrashy" so 
longest_consec(strarr, 2) should return "folingtrashy".

In the same way:
longest_consec(["zone", "abigail", "theta", "form", "libe", "zas", "theta", "abigail"], 2) --> "abigailtheta"

n being the length of the string array, if n = 0 or k > n or k <= 0 return "" (return Nothing in Elm, "nothing" in Erlang).

Note
consecutive strings : follow one after another without an interruption
*/

// my 
package kata

func LongestConsec(strarr []string, k int) string {
  if len(strarr) == 0 || k > len(strarr) || k <= 0  {
    return ""
  }
  
  longestString := ""
  maxLength := 0
  
  for i:= 0; i < len(strarr); i++ {
    if(i < (len(strarr)-k+1)) {
      tString := strarr[i]
      for j:= 1; j < k; j++ {
        tString = tString + strarr[i+j]
      }
      
      if len(tString) > maxLength {
          longestString = tString
          maxLength = len(tString)
      }
    } else {
      break
    }
  }
  
  return longestString
}

// bp
package kata

import ("strings")

func LongestConsec(strarr []string, k int) string {
	var buffer string
	var largest string

	for i := 0; i <= len(strarr)-k; i++ {
		buffer = strings.Join(strarr[i : i+k][:], "")
		if len(buffer) > len(largest) {
			largest = buffer
		}
	}
	return largest
}