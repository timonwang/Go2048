package main

import (
	"fmt"
	"math/rand"
)

func getNotNull(array []int) []int {
	result := make([]int, len(array))
	var j = 0
	for i := 0; i < len(array); i++ {
		if array[i] != 0 {
			result[j] = array[i]
			j = j + 1
		}
	}
	return result
}

func addSameInt(array []int) []int {
	var j = 0
	for i := 0; i < len(array); i++ {
		if array[i] != 0 && j == 0 {
			j = array[i]
		} else if array[i] == j {
			array[i-1] = array[i] + j
			for j := i; j < len(array)-1; j++ {
				array[j] = array[j+1]
			}
			array[len(array)-1] = 0
			break
		} else {
			j = array[i]
		}
	}
	return array
}

func reverse(array []int) []int {
	arrayLen := len(array)
	for i := 0; i < arrayLen/2; i++ {
		temp := array[i]
		array[i] = array[arrayLen-1-i]
		array[arrayLen-1-i] = temp
	}
	return array
}

func replace(matrix [][]int, index int, rowMode int, array []int) [][]int {
	if len(array) == len(matrix[0]) {
		if rowMode == 1 {
			for i := 0; i < len(matrix[0]); i++ {
				matrix[index][i] = array[i]
			}
		} else {
			for i := 0; i < len(matrix[0]); i++ {
				matrix[i][index] = array[i]
			}

		}
	}
	return matrix
}

func printMatrix(matrix [][]int) {
	matrixSize := len(matrix[0])
	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			fmt.Printf("%8d ", matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	var array = []int{0, 0, 1, 1}
	var matrix = [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	printMatrix(replace(matrix, 2, 0, array))
	fmt.Println(array)
	fmt.Println(addSameInt(getNotNull(array)))
	fmt.Println(reverse(addSameInt(getNotNull(array))))
	rand.Seed(1)
}
