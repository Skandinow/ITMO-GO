package mark

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"
	"unicode"
)

type Student struct {
	Name string
	Mark int
}

func NewStudent(name string, mark int) *Student {
	return &Student{Name: name, Mark: mark}
}

type StudentInfo struct {
	NumberOfGrades int
	SumOfGrades    int
}

func NewStudentInfo(numberOfGrades int, sumOfGrades int) *StudentInfo {
	return &StudentInfo{NumberOfGrades: numberOfGrades, SumOfGrades: sumOfGrades}
}

type StudentsStatistic interface {
	SummaryByStudent(student string) (int, bool)     // default_value, false - если студента нет
	AverageByStudent(student string) (float32, bool) // default_value, false - если студента нет
	Students() []string
	Summary() int
	Median() int
	MostFrequent() int
}

func NewStudentStatistic(studentsNameToStat map[string]StudentInfo, grades [10]int) *StudentStat {
	return &StudentStat{StudentsNameToStat: studentsNameToStat, Grades: grades}
}

const lenOfGrades = 10

type StudentStat struct {
	StudentsNameToStat map[string]StudentInfo // Вся сумма, количество
	Grades             [lenOfGrades]int
}

func (r *StudentStat) SummaryByStudent(student string) (int, bool) {
	studentInfo, isFound := r.StudentsNameToStat[student]
	if !isFound {
		return 0, false
	}

	return studentInfo.SumOfGrades, true
}
func average(studentInfo StudentInfo) float32 {
	ratio := math.Pow(10, 2.0)
	ratio32 := float32(ratio)
	round32 := float32(math.Round(float64(studentInfo.SumOfGrades) / float64(studentInfo.NumberOfGrades) * ratio))
	return round32 / ratio32

}
func (r *StudentStat) AverageByStudent(student string) (float32, bool) {
	studentInfo, isFound := r.StudentsNameToStat[student]
	if !isFound {
		return 0, false
	}

	return average(studentInfo), true
}

func (r *StudentStat) Students() []string {
	var result []string
	for key := range r.StudentsNameToStat {
		result = append(result, key)
	}

	sort.Slice(result,
		func(i, j int) bool {
			return r.StudentsNameToStat[result[i]].SumOfGrades > r.StudentsNameToStat[result[j]].SumOfGrades
		})

	return result
}

func (r *StudentStat) Summary() int {
	sum := 0
	for _, val := range r.StudentsNameToStat {
		sum += val.SumOfGrades
	}

	return sum
}

func (r *StudentStat) Median() int {
	sum := 0
	numb := 0
	for _, val := range r.StudentsNameToStat {
		sum += val.SumOfGrades
		numb += val.NumberOfGrades
	}
	if sum%numb == 0 {
		return sum / numb
	}

	return (sum / numb) + 1
}

func (r *StudentStat) MostFrequent() int {
	mstFreq := 0
	mstFreqInd := 0
	for i, element := range r.Grades {
		if element >= mstFreq {
			mstFreq = element
			mstFreqInd = i
		}
	}

	return mstFreqInd + 1
}
func someParser(grades [lenOfGrades]int, scanner *bufio.Scanner, studentsNameToStat map[string]StudentInfo) ([lenOfGrades]int, error) {
	for scanner.Scan() {
		name := ""
		number := ""
		wasSpecialChar := false
		isPrevDigit := false
		isBadString := false
		for _, element := range scanner.Text() {
			if element == '\t' || element == '\r' || unicode.IsDigit(element) {
				if element == '\t' || element == '\r' {
					wasSpecialChar = true
				}
				if unicode.IsDigit(element) && wasSpecialChar {
					number += string(element)
					isPrevDigit = true
				} else if unicode.IsDigit(element) {
					name += string(element)
				}

			} else {
				if isPrevDigit || wasSpecialChar {
					isBadString = true
					break
				}
				name += string(element)
			}
		}
		if isBadString || len(number) == 0 {
			continue
		}

		mark, err := strconv.Atoi(number)
		if mark < 0 || mark > 10 {
			continue
		}
		if err != nil {
			return grades, err
		}

		val, isFound := studentsNameToStat[name]

		if !isFound {
			studentInfo := NewStudentInfo(1, mark)
			studentsNameToStat[name] = *studentInfo
		} else {
			studentInfo := NewStudentInfo(val.NumberOfGrades+1, val.SumOfGrades+mark)
			studentsNameToStat[name] = *studentInfo
		}
		grades[mark-1] += 1
	}
	return grades, nil
}

func ReadStudentsStatistic(reader io.Reader) (StudentsStatistic, error) {
	scanner := bufio.NewScanner(reader)
	studentsNameToStat := map[string]StudentInfo{}
	var grades [lenOfGrades]int

	grades, err := someParser(grades, scanner, studentsNameToStat)
	if err != nil {
		return nil, err
	} else if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	studentStat := NewStudentStatistic(studentsNameToStat, grades)

	return studentStat, nil
}

func WriteStudentsStatistic(writer io.Writer, statistic StudentsStatistic) error {
	statTitleFormat := "%d\t%d\t%d\n"
	firstStr := fmt.Sprintf(statTitleFormat, statistic.Summary(), statistic.Median(), statistic.MostFrequent())
	_, err := writer.Write([]byte(firstStr))
	if err != nil {
		return err
	}

	stat := statistic.Students()

	sort.SliceStable(
		stat,
		func(i, j int) bool {
			first, _ := statistic.SummaryByStudent(stat[i])
			second, _ := statistic.SummaryByStudent(stat[j])
			return second < first
		},
	)

	for i, element := range stat {
		studentSum, _ := statistic.SummaryByStudent(element)
		average, _ := statistic.AverageByStudent(element)
		intAverage := int(average * 100)
		var iStr string
		n := ""
		if i != len(stat)-1 {
			n = "\n"
		}
		if intAverage%100 == 0 {
			studentWOFloatAccuracy := "%s\t%d\t%.0f%s"
			iStr = fmt.Sprintf(studentWOFloatAccuracy, element, studentSum, average, n)

		} else {
			studentWFloatAccuracy := "%s\t%d\t%.2f%s"
			iStr = fmt.Sprintf(studentWFloatAccuracy, element, studentSum, average, n)
		}
		_, err := writer.Write([]byte(iStr))
		if err != nil {
			return err
		}
	}

	return nil
}
