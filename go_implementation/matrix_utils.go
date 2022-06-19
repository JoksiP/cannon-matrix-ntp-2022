package main

import "math/rand"
import "fmt"

func createMatrix(N int, min, max float64) [][]float64{
    matrix := make([][]float64, N)
    for i := range matrix {
        matrix[i] = make([]float64, N)
        for j := range matrix[i]{
            matrix[i][j] = min + rand.Float64() * (max - min)
        }
    }
    return matrix
}

func createZeroMatrix(N int) [][]float64{
    matrix := make([][]float64, N)
    for i := range matrix {
        matrix[i] = make([]float64, N)
        for j := range matrix[i]{
            matrix[i][j] = 0.0
        }
    }
    return matrix
}

func shiftRows(matrix [][]float64, step_type string, k_proc int) [][]float64{
    N:=len(matrix)
    retMatrix := make([][]float64, N)  
    for i:= range retMatrix {
        retMatrix[i] = make([]float64, N)
        for j:= range retMatrix[i] {
            step:=1
            if step_type=="all"{
                step = i
            }
            if step_type=="k_proc"{
                step = k_proc
            }
            retMatrix[i][j] = matrix[i][(j+step) % N]
        }
    }
    return retMatrix
}

func shiftColumns(matrix [][]float64, step_type string, k_proc int) [][]float64{
    N:=len(matrix)
    retMatrix := make([][]float64, N)
    for i:=range retMatrix {
        retMatrix[i] = make([]float64, N)
        for j:=range retMatrix[i] {
            step:=1
            if step_type=="all"{
                step = j
            }
            if step_type=="k_proc"{
                step = k_proc
            }
            retMatrix[i][j] = matrix[(i+step) % N][j]
        }
    } 
    return retMatrix
}

func printMatrix(matrix1 [][]float64){
    N:=len(matrix1)
    for i:=0; i<N; i++ {
        for j:=0; j<N; j++ {
            fmt.Printf(" %f", matrix1[i][j])
        }
        fmt.Println()
        fmt.Println()
    }
}