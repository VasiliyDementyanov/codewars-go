/*
Find the missing letter
Write a method that takes an array of consecutive (increasing) letters as input and that returns the missing letter in the array.

You will always get an valid array.
And it will be always exactly one letter be missing.
The length of the array will always be at least 2.
The array will always contain letters in only one case.

Example:

['a','b','c','d','f'] -> 'e'
['O','Q','R','S'] -> 'P'
(Use the English alphabet with 26 letters!)

Have fun coding it and please don't forget to vote and rank this kata! :-)

I have also created other katas. Take a look if you enjoyed this kata!
*/

// my
package kata

func FindMissingLetter(chars []rune) rune {
	lastIndex := len(chars) - 1
	var resultRune rune
	for i := 0; i < lastIndex; i++ {
		if chars[i]+1 != chars[i+1] {
			resultRune = chars[i] + 1
			break
		}
	}

	return resultRune
}

// bp
package kata

func FindMissingLetter(a []rune) rune {
  c := a[0]
  for _,v :=range a[1:] {
    if  c++; v != c {break}
  }
  return c
}
