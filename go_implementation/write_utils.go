package main

import (
	"encoding/csv"
	"log"
	"os"
	"fmt"
	"strconv"
)

func matrix2String(matrix [][]float64) string{
    s := ""

   for i := range matrix {
      for _,j := range matrix[i] {
          s += fmt.Sprintf("%f ", j)
      }
	  s+=fmt.Sprintln()
   }

    return s
}

func initSeqWriteToCSV(){
	header := []string{"Iteration", "Matrix1", "Matrix2", "Result"}

	csvFile, err := os.Create("seq_outputs.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)
	
	csvwriter.Write(header)
	csvwriter.Flush()
	csvFile.Close()

}

func seqWriteToCSV(iter int, matrix1 , matrix2 , result [][]float64){
	csvFile, err := os.OpenFile("seq_outputs.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	data := []string{strconv.Itoa(iter), matrix2String(matrix1), matrix2String(matrix2), matrix2String(result)}
	fmt.Println(data[0])
	csvwriter := csv.NewWriter(csvFile)
	
	csvwriter.Write(data)
	csvwriter.Flush()
	csvFile.Close()
}