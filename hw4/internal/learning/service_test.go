package learning

import (
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"slices"
	"testing"
)

type testCase struct {
	name        string
	studentId   int64
	student     map[int64]studentInfo
	personalIds []int64
	groupIds    []int64
	result      []int64
	resultBool  bool
}

func initTestCases() []testCase {
	testCases := []testCase{
		{
			name:        "Tutor is found 1",
			studentId:   1,
			student:     map[int64]studentInfo{1: {"Petya", 16, "math"}},
			personalIds: []int64{1, 2, 3},
			groupIds:    []int64{5, 6, 7},
			result:      []int64{1, 2, 3},
			resultBool:  true,
		}, {
			name:        "Tutor is found 2",
			studentId:   2,
			student:     map[int64]studentInfo{2: {"Vasya", 16, "biology"}},
			personalIds: []int64{5, 6, 7},
			groupIds:    []int64{1, 2, 3},
			result:      []int64{5, 6, 7},
			resultBool:  true,
		}, {
			name:        "Group ids exist 1",
			studentId:   3,
			student:     map[int64]studentInfo{3: {"Valya", 17, "PE"}},
			personalIds: []int64{},
			groupIds:    []int64{10, 12},
			result:      []int64{10, 12},
			resultBool:  true,
		}, {
			name:        "Nothing exist 1",
			studentId:   4,
			student:     map[int64]studentInfo{4: {"Venya", 17, "Air Dynamics"}},
			personalIds: []int64{},
			groupIds:    []int64{},
			result:      nil,
		},
	}
	return testCases
}

func TestService_GetTutorsIDPreferIndividual(t *testing.T) {
	testCases := initTestCases()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			mc := minimock.NewController(t)
			repoMock := NewRepositoryMock(mc)
			individualService := NewIndividualServiceMock(t)
			groupService := NewGroupServiceMock(t)

			defer mc.Finish()

			student := tc.student[tc.studentId]
			repoMock.GetStudentInfoMock.Expect(tc.studentId).Return(&student, tc.resultBool)

			individualService.TutorsIDMock.Expect(student.Subject).Return(tc.result)

			groupService.TutorsIDMock.Expect(student.Subject).Return(tc.result)

			service := NewService(individualService, groupService, repoMock)
			result, isGood := service.GetTutorsIDPreferIndividual(tc.studentId)

			require.Equal(t, result, tc.result)
			require.Equal(t, isGood, tc.resultBool)

		})

	}

}

type subjectTestCase struct {
	name     string
	subjects []string
	topN     int
	isGood   bool
}

func TestService_GetTopSubjects(t *testing.T) {
	tutors := []int64{1, 2, 3, 4, 5}
	testCases := []subjectTestCase{
		{"test1", []string{"math", "biology", "pe", "astronomy", "physics"}, 3, true},
		{"test2", []string{"math", "biology", "pe"}, 2, true},
		{"test3", []string{"pe", "astronomy", "math", "biology"}, 5, false},
		{"test4", []string{"pe", "astronomy", "math", "biology"}, 3, true},
		{"test5", []string{"pe", "astronomy", "math"}, 3, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mc := minimock.NewController(t)
			repoMock := NewRepositoryMock(mc)
			individualService := NewIndividualServiceMock(t)
			groupService := NewGroupServiceMock(t)

			defer mc.Finish()

			var result []subjectInfo
			if tc.isGood {
				result = make([]subjectInfo, tc.topN)
				fillResult(tc, tutors, result)
				slices.Reverse(result)
				repoMock.GetAllSubjectsInfoMock.Expect().Return(result, tc.isGood)
			} else {
				result = nil
				repoMock.GetAllSubjectsInfoMock.Expect().Return(result, tc.isGood)

			}
			var resultStr []string
			if tc.isGood {
				resultStr = make([]string, tc.topN)
				for i := tc.topN - 1; i >= 0; i-- {
					resultStr[i] = tc.subjects[i]
				}
			}

			service := NewService(individualService, groupService, repoMock)
			res, isGood := service.GetTopSubjects(tc.topN)

			require.Equal(t, res, resultStr)
			require.Equal(t, isGood, tc.isGood)

		})

	}
}

func fillResult(tc subjectTestCase, tutors []int64, result []subjectInfo) {
	for i := tc.topN - 1; i >= 0; i-- {
		result[i].name = tc.subjects[i]
		result[i].numberOfTutors = tutors[i]
	}
}
