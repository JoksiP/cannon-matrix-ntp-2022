package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func matrix2String(matrix [][]float64) string {
	s := ""

	for i := range matrix {
		for _, j := range matrix[i] {
			s += fmt.Sprintf("%f ", j)
		}
		s += fmt.Sprintln()
	}

	return s
}

func initWriteToCSV(path string) {
	header := []string{"Iteration", "Matrix1", "Matrix2", "Result"}
	if path == "conc_outputs.csv" {
		header[0] = "Process"
	}
	csvFile, err := os.Create(path)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	csvwriter.Write(header)
	csvwriter.Flush()
	csvFile.Close()

}

func writeToCSV(iter int, matrix1, matrix2, result [][]float64, path string) {
	csvFile, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	data := []string{strconv.Itoa(iter), matrix2String(matrix1), matrix2String(matrix2), matrix2String(result)}
	csvwriter := csv.NewWriter(csvFile)

	csvwriter.Write(data)
	csvwriter.Flush()
	csvFile.Close()
}
