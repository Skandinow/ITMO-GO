package main

import (
	"fmt"
	"log"
	"os"

	"hw7/internal/fact"
)

func main() {
	slice := []int{-321, 100, -99, 98, 97, 96, 95, 94, 32}
	fmt.Println("Hello, world!")
	if err := fact.NewFactorization().Work(fact.Input{3, slice}, os.Stdout); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Finished")
}
