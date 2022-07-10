package main

import (
	//	"fmt"

	"math"
	"time"
)

type ConcThreadStruct struct {
	N      int
	Matrix [][]float64
}

func concThread(N int, matrix_thread [][][]float64, threads map[string]chan []float64, num int, master chan ConcThreadStruct) {
	n_p_sqrt := len(matrix_thread[0][0])
	for t := 0; t < N; t++ {
		matMul(matrix_thread[2], matrix_thread[0], matrix_thread[1], n_p_sqrt)
		if t == N-1 {
			master <- ConcThreadStruct{num, matrix_thread[2]}
			break
		}
		d1 := make([]float64, n_p_sqrt)
		for i := 0; i < n_p_sqrt; i++ {
			d1[i] = matrix_thread[0][i][0]
		}
		threads["dest1"] <- d1
		threads["dest2"] <- matrix_thread[1][0]
		s1 := <-threads["source1"]
		for i := 0; i < n_p_sqrt; i++ {
			matrix_thread[0][i] = append(matrix_thread[0][i][1:], s1[i])

		}
		s2 := <-threads["source2"]
		matrix_thread[1] = append(matrix_thread[1][1:], s2)

	}
}

func concImpl(matrix1 [][]float64, matrix2 [][]float64, P int) time.Duration {
	initWriteToCSV("conc_outputs.csv")
	N := len(matrix1)
	start := time.Now()
	//matrix1 = shiftRows(matrix1, step_type, 0)
	//matrix2 = shiftColumns(matrix2, step_type, 0)
	p_sqrt := int(math.Sqrt(float64(P)))
	n_p_sqrt := int(N / p_sqrt)
	result := createZeroMatrix(N)

	threads := make([][][]chan []float64, p_sqrt)
	master := make(chan ConcThreadStruct)

	for i := range threads {
		threads[i] = make([][]chan []float64, p_sqrt)
		for j := range threads[i] {
			threads[i][j] = make([]chan []float64, 2)
			for k := range threads[i][j] {
				threads[i][j][k] = make(chan []float64, 1)
			}
		}
	}

	for i := 0; i < p_sqrt; i++ {
		for j := 0; j < p_sqrt; j++ {
			matrix_thread := make([][][]float64, 3)
			t := make(map[string]chan []float64)
			matrix_thread[0] = make([][]float64, n_p_sqrt)
			matrix_thread[1] = make([][]float64, n_p_sqrt)
			matrix_thread[2] = make([][]float64, n_p_sqrt)
			row := i * n_p_sqrt
			col := j * n_p_sqrt
			for r := row; r < row+n_p_sqrt; r++ {
				r_n_p_sqrt := r % n_p_sqrt
				matrix_thread[0][r_n_p_sqrt] = make([]float64, n_p_sqrt)
				matrix_thread[1][r_n_p_sqrt] = make([]float64, n_p_sqrt)
				matrix_thread[2][r_n_p_sqrt] = make([]float64, n_p_sqrt)
				for c := col; c < col+n_p_sqrt; c++ {
					c_n_p_sqrt := c % n_p_sqrt
					matrix_thread[0][r_n_p_sqrt][c_n_p_sqrt] = matrix1[r][(c+r)%N]
					matrix_thread[1][r_n_p_sqrt][c_n_p_sqrt] = matrix2[(r+c)%N][c]
				}
			}
			t["dest1"] = threads[i][(j-1+p_sqrt)%p_sqrt][0]
			t["dest2"] = threads[(i-1+p_sqrt)%p_sqrt][j][1]
			t["source1"] = threads[i][j][0]
			t["source2"] = threads[i][j][1]
			go concThread(N, matrix_thread, t, i*p_sqrt+j, master)
		}
	}

	for t := 0; t < P; t++ {
		m := <-master
		for k := 0; k < n_p_sqrt; k++ {
			i := m.N/p_sqrt*n_p_sqrt + k
			j := m.N % p_sqrt * n_p_sqrt
			result[i] = append(append(result[i][:j], m.Matrix[k]...), result[i][j+n_p_sqrt:]...)
		}
	}
	//printMatrix(result)
	return time.Since(start)
}
