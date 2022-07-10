package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	N, _ := strconv.Atoi(os.Args[1])
	P, _ := strconv.Atoi(os.Args[2])

	p_sqrt := int(math.Sqrt(float64(P)))
	if p_sqrt*p_sqrt != P {
		fmt.Println("P must be a perfect square number.")
		os.Exit(0)
	}
	n_p_sqrt := float64(N) / float64(p_sqrt)
	if n_p_sqrt != float64(int(n_p_sqrt)) {
		fmt.Println("N must be dividable by square root of P")
		os.Exit(0)
	}

	matrix1 := createMatrix(N, 0.0, 10.0)
	matrix2 := createMatrix(N, 0.0, 10.0)
	//printMatrix(matrix1)
	//fmt.Println()
	//printMatrix(matrix2)
	//matMul(matrix1, matrix2)
	fmt.Println()
	fmt.Println()
	fmt.Printf("%s", seqImpl(matrix1, matrix2, N))
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Printf("%s", concImpl(matrix1, matrix2, P))
	//concImpl(matrix1, matrix2, 16)
}
