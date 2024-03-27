package main

import (
	"math"
	"math/cmplx"
)

func getCharByIndex(str string, idx int) rune {
	return []rune(str)[idx]
}

func getStringBySliceOfIndexes(str string, indexes []int) string {
	result := make([]rune, len(indexes))
	runes := []rune(str)
	for i, ind := range indexes {
		if ind >= 0 && ind < len(str) {
			result[i] = runes[ind]
		}
	}
	return string(result)
}

// как я понимаю, я складываю не указатели, а значения указателей
func addPointers(ptr1, ptr2 *int) *int {
	if ptr1 == nil || ptr2 == nil {
		return nil
	}
	result := *ptr1 + *ptr2
	return &result
}

func isComplexEqual(a, b complex128) bool {
	return math.Abs(real(a)-real(b)) < 1e-6 && math.Abs(imag(a)-imag(b)) < 1e6
}

func getRootsOfQuadraticEquation(a, b, c float64) (complex128, complex128) {
	d := b*b - 4*a*c
	x1 := (complex(-b, 0) + cmplx.Sqrt(complex(d, 0))) / complex(2*a, 0)
	x2 := (complex(-b, 0) - cmplx.Sqrt(complex(d, 0))) / complex(2*a, 0)

	return x1, x2
}

func mergeSort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	first := mergeSort(s[:len(s)/2])
	second := mergeSort(s[len(s)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	var final []int
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

func reverseSliceOne(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseSliceTwo(s []int) []int {
	ints := make([]int, len(s))

	copy(ints, s)

	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
		ints[i], ints[j] = ints[j], ints[i]
	}
	return ints
}

func swapPointers(a, b *int) {
	*a, *b = *b, *a

}

func isSliceEqual(a, b []int) bool {
	if len(a) > len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func deleteByIndex(s []int, idx int) []int {
	return append(s[:idx], s[idx+1:]...)
}

func main() {
}
