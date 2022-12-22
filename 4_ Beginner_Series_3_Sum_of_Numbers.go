/*
Given two integers a and b, which can be positive or negative,
find the sum of all the integers between and including them and return it. If the two numbers are equal return a or b.
Note: a and b are not ordered!
Examples (a, b) --> output (explanation)

(1, 0) --> 1 (1 + 0 = 1)
(1, 2) --> 3 (1 + 2 = 3)
(0, 1) --> 1 (0 + 1 = 1)
(1, 1) --> 1 (1 since both are same)
(-1, 0) --> -1 (-1 + 0 = -1)
(-1, 2) --> 2 (-1 + 0 + 1 + 2 = 2)
*/

// my
package kata

func GetSum(a, b int) int {
	if a == b {
		return a
	}

	if a > b {
		b, a = a, b
	}

	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}

	return sum
}

// bp
package kata

func GetSum(a, b int) int {
  if a > b {
    a, b = b, a
  }
  return (a + b) * (b - a + 1) / 2
}