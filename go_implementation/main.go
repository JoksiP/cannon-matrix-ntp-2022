package main

import (
	"fmt"
)

func main() {
	N := 256
	P := 16
	matrix1 := createMatrix(N, 0.0, 10.0)
	matrix2 := createMatrix(N, 0.0, 10.0)
	//printMatrix(matrix1)
	//fmt.Println()
	//printMatrix(matrix2)
	//matMul(matrix1, matrix2)
	fmt.Println()
	fmt.Println()
	fmt.Printf("%s", seqImpl(matrix1, matrix2, N))
	fmt.Printf("%s", concImpl(matrix1, matrix2, P))
	//concImpl(matrix1, matrix2, 16)
}
