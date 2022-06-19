package main

import "time"

func seqImpl(matrix1 [][]float64, matrix2 [][]float64, P int) time.Duration {
	initWriteToCSV("seq_outputs.csv")
	step_type := "all"
	start := time.Now()

	matrix1 = shiftRows(matrix1, step_type, 0)
	matrix2 = shiftColumns(matrix2, step_type, 0)
	N := len(matrix1)
	//printMatrix(result)
	//printMatrix(matrix2)
	result := createZeroMatrix(N)
	for k := 0; k < P; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				result[i][j] += matrix1[i][j] * matrix2[i][j]
			}
		}
		step_type := "single"
		matrix1 = shiftRows(matrix1, step_type, 0)
		matrix2 = shiftColumns(matrix2, step_type, 0)
		//writeToCSV(k, matrix1, matrix2, result, "seq_outputs.csv")
	}
	//printMatrix(result)
	return time.Since(start)
}
