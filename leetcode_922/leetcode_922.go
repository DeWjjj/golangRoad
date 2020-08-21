package main

import (
	"fmt"
)

func sortArrayByParityII(A []int) []int {
	var arrayA []int
	var arrayB []int
	var arrayC []int

	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			arrayA = append(arrayA, A[i])
		} else {
			arrayB = append(arrayB, A[i])
		}
	}
	var (
		indexA = 0
		indexB = 0
		indexC = 0
	)

	for indexC != len(A) {
		if indexC%2 == 0 {
			arrayC = append(arrayC, arrayA[indexA])
			indexC++
			indexA++
		} else {
			arrayC = append(arrayC, arrayB[indexB])
			indexC++
			indexB++
		}

	}
	return arrayC
}

func main() {
	var A = []int{4, 2, 5, 7}
	value := sortArrayByParityII(A)
	fmt.Println(value)
}
