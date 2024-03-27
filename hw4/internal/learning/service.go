package learning

import (
	"sort"
)

type ServiceInterface interface {
	GetTutorsIDPreferIndividual(studentID int64) ([]int64, bool)
	GetTopSubjects(topN int) ([]string, bool)
}

type RepositoryInterface interface {
	GetStudentInfo(id int64) (*studentInfo, bool)
	GetAllSubjects() ([]string, bool)
	GetAllSubjectsInfo() ([]subjectInfo, bool)
}

type TutorServiceInterface interface {
	Subjects() []string
	TutorsID(subject string) []int64
}

type Service struct {
	individualTutorService TutorServiceInterface
	groupTutorService      TutorServiceInterface
	repository             RepositoryInterface
}

func NewService(
	individualTutorService TutorServiceInterface,
	groupTutorService TutorServiceInterface,
	repository RepositoryInterface,
) *Service {
	return &Service{
		individualTutorService: individualTutorService,
		groupTutorService:      groupTutorService,
		repository:             repository,
	}
}

func (s *Service) GetTutorsIDPreferIndividual(studentID int64) ([]int64, bool) {
	studentInfo, ok := s.repository.GetStudentInfo(studentID)
	if !ok {
		return nil, ok
	}

	tutorsID := s.individualTutorService.TutorsID(studentInfo.Subject)
	if len(tutorsID) == 0 {
		tutorsID = s.groupTutorService.TutorsID(studentInfo.Subject)
		if len(tutorsID) == 0 {
			return nil, false
		}
		return tutorsID, true
	}

	return tutorsID, true
}

func (s *Service) GetTopSubjects(topN int) ([]string, bool) {
	subjects, ok := s.repository.GetAllSubjectsInfo()
	if !ok {
		return nil, ok
	}
	if len(subjects) > topN {
		return nil, false
	}

	sort.SliceStable(
		subjects,
		func(i, j int) bool {
			return subjects[i].numberOfTutors < subjects[j].numberOfTutors
		},
	)

	return fromSubject(subjects[:topN]), true
}
