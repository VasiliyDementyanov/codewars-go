/*
Write an algorithm that takes an array and moves all of the zeros to the end, preserving the order of the other elements.
MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}) //
returns []int{ 1, 2, 1, 1, 3, 1, 0, 0, 0, 0 }
*/

// my
package kata

func MoveZeros(arr []int) []int {
	var resultArr []int
	var zeroArr []int

	for _, v := range arr {
		if v != 0 {
			resultArr = append(resultArr, v)
		} else {
			zeroArr = append(zeroArr, v)
		}
	}

	resultArr = append(resultArr, zeroArr...)
	return resultArr
}

// bp
package kata

func MoveZeros(arr []int) []int {
 	res:= make([]int,len(arr))
	ind:=0
	for i:=0;i<len(arr);i++{
		if arr[i]!=0{
			res[ind]=arr[i]
			ind++
		}
	}
	return res
}