package rangeI

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewRangeInt(t *testing.T) {
	range1 := NewRangeInt(1, 10)
	range2 := NewRangeInt(9, 5)

	require.NotNil(t, range1)
	require.NotNil(t, range2)
}

func TestRange_Length(t *testing.T) {
	range1 := NewRangeInt(1, 10)
	range2 := NewRangeInt(9, 5)

	require.Equal(t, range1.Length(), 10)
	require.Equal(t, range2.Length(), 0)
}

func TestRange_ContainsInt(t *testing.T) {
	range1 := NewRangeInt(3, 8)

	for i := 0; i < 10; i++ {
		if i >= 3 && i <= 8 {
			require.True(t, range1.ContainsInt(i))
		} else {
			require.False(t, range1.ContainsInt(i))
		}
	}
}

func TestRange_ContainsRange(t *testing.T) {
	range1 := NewRangeInt(3, 8)

	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			if j >= 3 && i <= 8 {
				require.True(t, range1.ContainsRange(NewRangeInt(j, i)))
			} else {
				require.False(t, range1.ContainsRange(NewRangeInt(j, i)))
			}
		}
	}
}

func TestRange_IsEmpty(t *testing.T) {
	range1 := NewRangeInt(3, 8)

	range2 := NewRangeInt(5, 12)
	range2.Intersect(NewRangeInt(0, 1))

	require.False(t, range1.IsEmpty())
	require.True(t, range2.IsEmpty())
}

func TestRange_Maximum(t *testing.T) {
	range1 := NewRangeInt(3, 8)
	range2 := NewRangeInt(5, 12)
	range3 := NewRangeInt(12, 5)

	max1, _ := range1.Maximum()
	max2, _ := range2.Maximum()
	max3, _ := range3.Maximum()

	require.Equal(t, max1, 8)
	require.Equal(t, max2, 12)
	require.Equal(t, max3, 0)
}

func TestRange_Minimum(t *testing.T) {
	range1 := NewRangeInt(3, 8)
	range2 := NewRangeInt(5, 12)
	range3 := NewRangeInt(12, 5)

	max1, _ := range1.Minimum()
	max2, _ := range2.Minimum()
	max3, _ := range3.Minimum()

	require.Equal(t, max1, 3)
	require.Equal(t, max2, 5)
	require.Equal(t, max3, 0)
}

func TestRange_Intersect(t *testing.T) {
	range1 := NewRangeInt(3, 8)
	range2 := NewRangeInt(5, 12)
	range3 := NewRangeInt(12, 5)

	range1.Intersect(range2)
	min1, _ := range1.Minimum()
	max1, _ := range1.Maximum()
	require.Equal(t, min1, 5)
	require.Equal(t, max1, 8)

	range1.Intersect(range3)
	min1, _ = range1.Minimum()
	max1, _ = range1.Maximum()
	require.Equal(t, min1, 0)
	require.Equal(t, max1, 0)

}

func TestRange_IsIntersect(t *testing.T) {
	range1 := NewRangeInt(3, 8)
	range2 := NewRangeInt(5, 12)
	range3 := NewRangeInt(12, 5)
	range4 := NewRangeInt(60, 62)

	require.True(t, range1.IsIntersect(range2))
	require.False(t, range1.IsIntersect(range3))
	require.False(t, range1.IsIntersect(range4))

}

func TestRange_Union(t *testing.T) {
	range1 := NewRangeInt(3, 8)
	range2 := NewRangeInt(5, 12)

	range1.Union(range2)
	min1, _ := range1.Minimum()
	max1, _ := range1.Maximum()
	require.Equal(t, min1, 3)
	require.Equal(t, max1, 12)
}

func TestRange_String(t *testing.T) {
	range1 := NewRangeInt(3, 8)
	range2 := NewRangeInt(5, 12)
	range3 := NewRangeInt(12, 5)

	require.Equal(t, range1.String(), "[3,8]")
	require.Equal(t, range2.String(), "[5,12]")
	require.Equal(t, range3.String(), "")
}

func TestRange_ToSlice(t *testing.T) {
	range1 := NewRangeInt(3, 8)
	range2 := NewRangeInt(5, 12)
	range3 := NewRangeInt(12, 5)

	require.Equal(t, range1.ToSlice(), []int{3, 4, 5, 6, 7, 8})
	require.Equal(t, range2.ToSlice(), []int{5, 6, 7, 8, 9, 10, 11, 12})
	require.Equal(t, range3.ToSlice(), []int{})
}
