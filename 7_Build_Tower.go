/*
Build Tower
Build a pyramid-shaped tower, as an array/list of strings, given a positive integer number of floors.
A tower block is represented with "*" character.

For example, a tower with 3 floors looks like this:
[
  "  *  ",
  " *** ",
  "*****"
]
And a tower with 6 floors looks like this:
[
  "     *     ",
  "    ***    ",
  "   *****   ",
  "  *******  ",
  " ********* ",
  "***********"
]
*/

// my
package kata

import "strings"

func TowerBuilder(nFloors int) []string {
	var resultTree []string

	floor := 2*nFloors - 1
	for i := 1; i <= nFloors; i++ {
		currentLevelStar := 2*i - 1
		currentLevelLeftRightEmpty := (floor - currentLevelStar) / 2

		currentLevelTree := strings.Repeat(" ", currentLevelLeftRightEmpty) + strings.Repeat("*", currentLevelStar) + strings.Repeat(" ", currentLevelLeftRightEmpty)
		resultTree = append(resultTree, currentLevelTree)
	}

	return resultTree
}

// bp
package kata

import "strings"

func TowerBuilder(nFloors int) []string {
  tower := make([]string, nFloors)
  for i := nFloors; i > 0; i-- {
    starsCount := i * 2 - 1
    spaceCount := (nFloors * 2 - 1 - starsCount) / 2
    tower[i-1] = strings.Repeat(" ", spaceCount) + strings.Repeat("*", starsCount) + strings.Repeat(" ", spaceCount)
  }
  return tower
}