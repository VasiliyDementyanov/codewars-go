/*
My friend John and I are members of the "Fat to Fit Club (FFC)".
John is worried because each month a list with the weights of
members is published and each month he is the last on the list which means he is the heaviest.
I am the one who establishes the list so I told him: "Don't worry any more,
I will modify the order of the list". It was decided to attribute a "weight" to numbers.
The weight of a number will be from now on the sum of its digits.

For example 99 will have "weight" 18, 100 will have "weight"
1 so in the list 100 will come before 99.

Given a string with the weights of
FFC members in normal order can you give this string ordered by "weights" of these numbers?

Example:
"56 65 74 100 99 68 86 180 90" ordered by numbers weights becomes:

"100 180 90 56 65 74 68 86 99"
When two numbers have the same "weight", let us class
them as if they were strings (alphabetical ordering) and not numbers:

180 is before 90 since, having the same "weight" (9), it comes before as a string.

All numbers in the list are positive numbers and the list can be empty.

Notes
it may happen that the input string have leading,
trailing whitespaces and more than a unique whitespace between two consecutive numbers
For C: The result is freed.
*/

// my
package kata

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func OrderWeight(strng string) string {
	re := regexp.MustCompile(`[0-9]+`)
	numsString := re.FindAllString(strng, -1)

	sort.Slice(numsString, func(i, j int) bool {
		sumI := 0
		sumJ := 0

		for _, char := range numsString[i] {
			n, _ := strconv.Atoi(string(char))
			sumI += n
		}

		for _, char := range numsString[j] {
			n, _ := strconv.Atoi(string(char))
			sumJ += n
		}

		if sumI == sumJ {
			for k := 0; len(numsString[i]) > k && len(numsString[j]) > k; k++ {
				if numsString[i][k] == numsString[j][k] {
					continue
				}
				return numsString[i][k] < numsString[j][k]
			}

			return len(numsString[i]) < len(numsString[j])
		}

		return sumI < sumJ
	})

	return strings.Join(numsString, " ")
}

// bp
package kata

import (
	"sort"
	"strings"
)

type byDigitSum []string

func digitSum(s string) int {
	sum := 0

	for _, rune := range s {
		sum += int(rune - '0')
	}

	return sum
}

func (ns byDigitSum) Len() int {
	return len(ns)
}

func (ns byDigitSum) Swap(i, j int) {
	ns[i], ns[j] = ns[j], ns[i]
}

func (ns byDigitSum) Less(i, j int) bool {
	ds1 := digitSum(ns[i])
	ds2 := digitSum(ns[j])

	return ds1 < ds2 || ds1 == ds2 && strings.Compare(ns[i], ns[j]) == -1
}

func OrderWeight(strng string) string {
	ns := strings.Split(strng, " ")
	sort.Sort(byDigitSum(ns))
	return strings.Join(ns, " ")
}