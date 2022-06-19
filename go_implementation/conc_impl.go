package main

import (
	//	"fmt"
	"fmt"
	"math"
	"time"
)

type ConcThreadStruct struct {
	matrix1  [][]float64
	matrix2  [][]float64
	result   [][]float64
	N        int
	p_step_i int
	p_step_j int
	n_p_sqrt int
}

type RetStruct struct {
	matrix <-chan [][]float64
}

func concThread(cts ConcThreadStruct, i int) <-chan [][]float64 {
	step_type := "single"
	r := make(chan [][]float64)
	go func() {
		defer close(r)
		for k := 0; k < cts.n_p_sqrt*cts.n_p_sqrt; k++ {
			for i := cts.p_step_i; i < cts.p_step_i+cts.n_p_sqrt; i++ {
				for j := cts.p_step_j; j < cts.p_step_j+cts.n_p_sqrt; j++ {
					cts.result[i][j] += cts.matrix1[i][j] * cts.matrix2[i][j]
				}
			}
			//shift
			cts.matrix1 = shiftRows(cts.matrix1, step_type, 0)
			cts.matrix2 = shiftColumns(cts.matrix2, step_type, 0)
		}
		//writeToCSV(i, cts.matrix1, cts.matrix2, cts.result, "conc_outputs.csv")
		r <- cts.result
	}()
	return r
}

func concImpl(matrix1 [][]float64, matrix2 [][]float64, P int) time.Duration {
	initWriteToCSV("conc_outputs.csv")
	N := len(matrix1)
	step_type := "all"
	matrix1 = shiftRows(matrix1, step_type, 0)
	matrix2 = shiftColumns(matrix2, step_type, 0)
	p_sqrt := int(math.Sqrt(float64(P)))
	n_p_sqrt := int(N / p_sqrt)
	result := createZeroMatrix(N)

	threads := make([]ConcThreadStruct, P)
	for i := 0; i < p_sqrt; i++ {
		for j := 0; j < p_sqrt; j++ {
			threads[i*p_sqrt+j] = ConcThreadStruct{matrix1, matrix2, createZeroMatrix(N), N, i * n_p_sqrt, j * n_p_sqrt, n_p_sqrt}
		}
	}

	//wg := &sync.WaitGroup{}
	//wg.Add(len(threads))
	start := time.Now()
	all_r := make([]RetStruct, P)
	for i, xi := range threads {
		//go func(i int, xi ConcThreadStruct) {
		//	defer wg.Done()
		all_r[i].matrix = concThread(xi, i)
		//}(i, xi)
	}
	for i := 0; i < P; i++ {
		addMatrix(result, <-all_r[i].matrix)
	}
	//wg.Wait()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	//printMatrix(result)
	return time.Since(start)
}
