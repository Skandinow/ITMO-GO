package main

import (
	"fmt"
	m "hw5/internal/mark"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./data/input_1.tsv")
	if err != nil {
		log.Fatal("error with reading input file: ", err)
	}
	studentsStatistic, err := m.ReadStudentsStatistic(file)
	if err != nil {
		log.Fatal("error with reading input file: ", err)
	}
	studentsStatistic.Students()

	f, err := os.Create("file_name.tcv")
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	err2 := m.WriteStudentsStatistic(f, studentsStatistic)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}
