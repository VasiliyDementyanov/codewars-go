/*
You are given an array (which will have a length of at least 3, but could be very large) containing integers.
The array is either entirely comprised of odd integers or entirely comprised of even integers except for a single integer N.
Write a method that takes the array as an argument and returns this "outlier" N.

[2, 4, 0, 100, 4, 11, 2602, 36]
Should return: 11 (the only odd number)

[160, 3, 1719, 19, 11, 13, -21]
Should return: 160 (the only even number)
*/

// my
package kata

func FindOutlier(integers []int) int {
	isOutlierEven := false
	outlier := 0

	var evens []int
	var odds []int

	for i := 0; i < 3; i++ {
		if integers[i]%2 == 0 {
			evens = append(evens, integers[i])
		} else {
			odds = append(odds, integers[i])
		}
	}

	if len(evens) == 2 {
		return odds[0]
	}

	if len(odds) == 2 {
		return evens[0]
	}

	if len(evens) == 3 {
		isOutlierEven = false
	}

	if len(odds) == 3 {
		isOutlierEven = true
	}

	for i := 3; i < len(integers); i++ {
		if integers[i]%2 == 0 {
			if isOutlierEven {
				outlier = integers[i]
				break
			}
		} else {
			if !isOutlierEven {
				outlier = integers[i]
				break
			}
		}
	}

	return outlier
}

// bp
package kata

func FindOutlier(integers []int) int {
	var odd, even *int
	for i, integer := range integers {
		if even != nil && odd != nil {
			if integer % 2 == 0 {
				return *odd
			}
			return *even
		}
		if integer % 2 == 0 {
			even = &integers[i]
		} else {
			odd = &integers[i]
		}
	}
	return integers[len(integers)-1]
}