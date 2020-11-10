package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	if _, err := os.Stat("contracts.xlsx"); os.IsNotExist(err) {
		fmt.Printf("File does not exists\n")
		os.Exit(1)
	}

	fmt.Printf("Now read the file\n")

	readExcel("contracts.xlsx")

	// lines, err := readLines("contracts.xlsx")
	// if err != nil {
	// 	log.Fatalf("readLines: %s", err)
	// }

	// for i, line := range lines {
	// 	fmt.Println(i, line)
	// }
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func readExcel(path string) {
	xlsx, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sheetA, err := xlsx.GetRows("A")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sheetB, err := xlsx.GetRows("B")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, row := range sheetA {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}

	for _, row := range sheetB {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
