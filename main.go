package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func getPoint() (x, y int) {
	rand.Seed(time.Now().Unix() % 13)
	x = rand.Int() % 4
	rand.Seed(time.Now().Unix() % 53)
	y = rand.Int() % 4
	return x, y
}

func getNumber() int {
	rand.Seed(time.Now().Unix() % 11)
	return (rand.Intn(2) + 1) * 2
}

func checkGameOver(matrix [4][4]int) bool {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if matrix[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func getDirection() byte {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadByte()
	return input
}

func getNotZero(array []int) []int {
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
		} else if array[i] == j && j != 0 {
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

func replace(matrix [4][4]int, index int, rowMode int, array []int) [4][4]int {
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

func doMerge(direct byte, matrix [4][4]int) [4][4]int {
	switch direct {
	case 'w':
		for j := 0; j < 4; j++ {
			merged := []int{0, 0, 0, 0}
			for i := 0; i < 4; i++ {
				merged[i] = matrix[i][j]
			}
			fmt.Println(merged)
			merged = addSameInt(getNotZero(merged))
			fmt.Println("+++++++++")
			fmt.Println(merged)
			matrix = replace(matrix, j, 0, merged)
		}
		break
	case 's':
		for j := 0; j < 4; j++ {
			merged := []int{0, 0, 0, 0}
			for i := 0; i < 4; i++ {
				merged[i] = matrix[i][j]
			}
			fmt.Println(merged)
			merged = addSameInt(getNotZero(reverse(merged)))
			fmt.Println("+++++++++")
			fmt.Println(merged)
			matrix = replace(matrix, j, 0, merged)
		}
		break
	case 'a':
		for i := 0; i < 4; i++ {
			merged := []int{0, 0, 0, 0}
			for j := 0; j < 4; j++ {
				merged[j] = matrix[i][j]
			}
			fmt.Println(merged)
			merged = addSameInt(getNotZero(merged))
			fmt.Println("+++++++++")
			fmt.Println(merged)
			matrix = replace(matrix, i, 1, merged)
		}
		break
	case 'd':
		for i := 0; i < 4; i++ {
			merged := []int{0, 0, 0, 0}
			for j := 0; j < 4; j++ {
				merged[j] = matrix[i][j]
			}
			fmt.Println(merged)
			merged = addSameInt(getNotZero(reverse(merged)))
			fmt.Println("+++++++++")
			fmt.Println(merged)
			matrix = replace(matrix, i, 1, merged)
		}
		break
	case 'x':
		break
	default:
		break
	}
	return matrix
}

func main() {
	var count = 0
	var gameMatrix = [4][4]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}

	for {
		var x, y = getPoint()
		if gameMatrix[x][y] > 0 {
			var i = 0
			var j = 0
			for i = 0; i >= 0 && i < 4; i++ {
				for j = 0; j >= 0 && j < 4; j++ {
					if gameMatrix[i][j] == 0 {
						x = i
						y = j
						break
					}
				}
			}
		}
		if gameMatrix[x][y] == 0 {
			gameMatrix[x][y] = getNumber()
			if count += 1; count >= 2 {
				if checkGameOver(gameMatrix) {
					fmt.Println("Game Over!!!!!! You Loss!!!")
					fmt.Printf("You use %d steps, Get %d marks\n", count, getTotalMarks(count, gameMatrix))
					break
				}
				printResult(gameMatrix)
			reDo:
				var direct = getDirection()
				if direct == 'x' {
					break
				} else if direct == 'w' || direct == 's' || direct == 'a' || direct == 'd' {
					gameMatrix = doMerge(direct, gameMatrix)
					fmt.Println("------------------------------------------")
					printResult(gameMatrix)
					fmt.Println("******************************************")
				} else {
					fmt.Println("Retry again ,your input is not valid.")
					fmt.Println("w -- up\n\ts -- down\n\ta -- left\n\td -- right\n\tx -- exit game")
					goto reDo
				}
			}
		}
	}
}

func getTotalMarks(count int, matrix [4][4]int) int {
	var result = 0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result += matrix[i][j]
		}
	}
	return result + count
}

func printResult(matrix [4][4]int) {
	fmt.Println("Your Result:")
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("%6d", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}
