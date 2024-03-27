package fact

import (
	"errors"
	"io"
	"sync"
)

type Input struct {
	NumsOfGoroutine int   // n - число горутин
	Numbers         []int // слайс чисел, которые необходимо факторизовать
}

type Factorization interface {
	Work(Input, io.Writer) error
}

type FactorizationImpl struct {
}

func NewFactorization() *FactorizationImpl {
	return &FactorizationImpl{}
}

func factorize(a int) {

}

func (r *FactorizationImpl) Work(input Input, writer io.Writer) error {
	n := input.NumsOfGoroutine
	size := len(input.Numbers)
	i := 0

	if n < 1 {
		return errors.New("There is no Goroutine for evaluation ")
	}

	var wg sync.WaitGroup

	for i <= size {
		for k := 0; k < n; k++ {
			wg.Add(1)
			go factorize(input.Numbers[i+k])
		}
		i += n
	}

	if i != 0 {

	}

	return nil
}
