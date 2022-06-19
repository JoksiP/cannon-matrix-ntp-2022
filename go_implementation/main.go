package main
//import "fmt"

func main() {
    N := 10
    matrix1 := createMatrix(N, 0.0, 10.0)
    matrix2 := createMatrix(N, 0.0, 10.0)
    //printMatrix(matrix1)
    //fmt.Println()
    //printMatrix(matrix2)
    seqImpl(matrix1, matrix2, N)
}

func seqImpl(matrix1 [][]float64, matrix2 [][]float64, P int) {
    initSeqWriteToCSV()
    step_type:="all"
    matrix1 = shiftRows(matrix1, step_type, 0)
    matrix2 = shiftColumns(matrix2, step_type, 0)
    N:=len(matrix1)
    //printMatrix(result)
    //printMatrix(matrix2)
    result := createZeroMatrix(N)
    for k:=0;k<P;k++ {
        for i:=0;i<N;i++{
            for j:=0;j<N;j++{
                result[i][j] += matrix1[i][j] * matrix2[i][j]
            }
        }
        step_type:="single"
        matrix1 = shiftRows(matrix1, step_type, 0)
        matrix2 = shiftColumns(matrix2, step_type, 0)
        seqWriteToCSV(k, matrix1, matrix2, result)
    }
    printMatrix(result)
}

