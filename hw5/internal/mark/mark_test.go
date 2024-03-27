package mark

import (
	"github.com/stretchr/testify/require"
	"math"
	"sort"
	"testing"
)

func initTestCase() *StudentStat {
	return NewStudentStatistic(
		map[string]StudentInfo{
			"George Gelmetdinov\t": {4, 29},
			"Andrew Hopkins":       {10, 55},
			"Egor Sviridenko":      {2, 10},
			"Dan Latanov\t":        {14, 89}},
		[10]int{3, 2, 2, 2, 2, 3, 4, 5, 4, 3},
	)
}

func TestStudentStat_Median(t *testing.T) {
	testCase := initTestCase()
	numberOfElements := 0
	sum := 0
	for i, element := range testCase.Grades {
		numberOfElements += element
		sum += (i + 1) * element
	}
	result := 0
	if sum%numberOfElements == 0 {
		result = sum / numberOfElements
	} else {
		result = sum/numberOfElements + 1
	}

	require.Equal(t, result, testCase.Median())
}

func TestStudentStat_Summary(t *testing.T) {
	testCase := initTestCase()
	sum := 0
	for i, element := range testCase.Grades {
		sum += (i + 1) * element
	}

	require.Equal(t, sum, testCase.Summary())
}

func TestStudentStat_Students(t *testing.T) {
	testCase := initTestCase()
	requiredStudents := testCase.Students()
	students := make([]string, 0, len(requiredStudents))
	for key := range testCase.StudentsNameToStat {
		students = append(students, key)
	}
	sort.Slice(
		students,
		func(i, j int) bool {
			return students[i] < students[j]
		},
	)

	sort.Slice(
		requiredStudents,
		func(i, j int) bool {
			return requiredStudents[i] < requiredStudents[j]
		},
	)

	require.Equal(t, students, requiredStudents)
}

func TestStudentStat_SummaryByStudent(t *testing.T) {
	testCase := initTestCase()

	for key, val := range testCase.StudentsNameToStat {
		res, _ := testCase.SummaryByStudent(key)
		require.Equal(t, val.SumOfGrades, res)
	}
}

func TestStudentStat_AverageByStudent(t *testing.T) {
	testCase := initTestCase()
	ratio := math.Pow(10, 2.0)
	ratio32 := float32(ratio)

	for key, val := range testCase.StudentsNameToStat {
		res, _ := testCase.AverageByStudent(key)
		round32 := float32(math.Round(float64(val.SumOfGrades) / float64(val.NumberOfGrades) * ratio))
		require.Equal(t, round32/ratio32, res)

	}
}
func TestStudentStat_MostFrequent(t *testing.T) {
	testCase := initTestCase()

	require.Equal(t, 8, testCase.MostFrequent())
}
