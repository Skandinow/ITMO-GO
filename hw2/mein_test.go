package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetCharByIndex2(t *testing.T) {
	str := "abobazavr"
	for i, elem := range str {
		require.Equal(t, getCharByIndex(str, i), elem)
	}
}

func TestGetStringBySliceOfIndexes2(t *testing.T) {
	indexes := []int{0, 1, 6, 7}
	str := "abobazavr"

	require.Equal(t, getStringBySliceOfIndexes(str, indexes), "abav")
}

func TestAddPointers2(t *testing.T) {
	var first, second = 1, 2
	require.Equal(t, *addPointers(&first, &second), 3)
}

func TestIsComplexEqual2(t *testing.T) {
	a := 1i + 2
	b := 1i + 5
	require.False(t, isComplexEqual(a, b))
}

func TestGetRootsOfQuadraticEquation2(t *testing.T) {
	var x float64 = 1
	var y float64 = 4
	var z float64 = 3
	x1, x2 := getRootsOfQuadraticEquation(x, y, z)
	require.True(t, isComplexEqual(x1, complex(-1, 0)))
	require.True(t, isComplexEqual(x2, complex(-3, 0)))
}

func TestMergeSort2(t *testing.T) {
	ar1 := []int{9, 2, 3, 1, 10}
	require.Equal(t, mergeSort(ar1), []int{1, 2, 3, 9, 10})
}

func TestReverseSliceOne2(t *testing.T) {
	indexes := []int{0, 1, 6, 7}

	reverseSliceOne(indexes)
	require.Equal(t, indexes, []int{7, 6, 1, 0})
}
func TestReverseSliceTwo2(t *testing.T) {
	indexes := []int{0, 1, 6, 7}

	require.Equal(t, reverseSliceTwo(indexes), []int{7, 6, 1, 0})
	require.Equal(t, indexes, []int{0, 1, 6, 7})

}

func TestIsSliceEqual2(t *testing.T) {
	ar1 := []int{1, 2, 3, 4, 5}
	ar2 := []int{1, 2, 3, 4, 4}
	require.False(t, isSliceEqual(ar1, ar2))
}
func TestDeleteByIndex2(t *testing.T) {
	ar1 := []int{1, 2, 3, 4, 5}
	ar2 := []int{1, 2, 4, 5}
	require.Equal(t, deleteByIndex(ar1, 2), ar2)
}
