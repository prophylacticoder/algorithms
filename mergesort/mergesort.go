package main

import (
	"fmt"
)

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	leftSide := mergeSort(arr[:len(arr)/2])
	rightSide := mergeSort(arr[len(arr)/2:])

	sortedArr := []int{}
	j, k := 0, 0
	for i := 0; j < len(leftSide) && k < len(rightSide); i++ {
		if k < len(rightSide) && j < len(leftSide) {
			if leftSide[j] < rightSide[k] {
				sortedArr = append(sortedArr, leftSide[j])
				j++
			} else if rightSide[k] < leftSide[j] {
				sortedArr = append(sortedArr, rightSide[k])
				k++
			} else if leftSide[j] == rightSide[k] {
				sortedArr = append(sortedArr, leftSide[j], rightSide[k])
				j++
				k++
				i++
			}
		}
	}

	if j < len(leftSide) {
		for ; j < len(leftSide); j++ {
			sortedArr = append(sortedArr, leftSide[j])
		}
	} else if k < len(rightSide) {
		for ; k < len(rightSide); k++ {
			sortedArr = append(sortedArr, rightSide[k])
		}
	}

	return sortedArr
}

func main() {
	arr := []int{5, 3, 1, 9, 0, 1, 1}
	arr = mergeSort(arr)
	fmt.Println(arr)
}
