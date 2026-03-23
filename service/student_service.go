package service

import (
	"errors"
	"math"
	"test/model"
	"test/repo"
)

type StudentServiceImpl struct {
	Repo repo.StudentRepository
}

func NewStudentService(repo repo.StudentRepository) *StudentServiceImpl {
	return &StudentServiceImpl{Repo: repo}
}

func (s *StudentServiceImpl) AddStudentGrade(student model.Student, reportCard model.ReportCard) (string, error) {
	if student.Name == "" {
		return "", errors.New("empty student name")
	}

	if reportCard.Grade < 0 || reportCard.Grade > 100 {
		return "", errors.New("invalid grade")
	}

	if reportCard.Subject == "" {
		return "", errors.New("subject is empty")
	}

	student, err := s.Repo.FindStudentById(student.Id)
	if err != nil {
		isSuccess, _ := s.Repo.AddStudent(student)
		if !isSuccess {
			return "", err
		}
	}

	isSuccess, _ := s.Repo.AddGrade(reportCard)
	if !isSuccess {
		return "", err
	}

	return "Add Success", nil
}

func (s *StudentServiceImpl) CalculateStudentAverage(studentId int) (float64, error) {
	var totalGrade float64
	reportCards, _ := s.Repo.FindGradesByStudentId(studentId)
	if len(reportCards) <= 0 {
		return 0.0, nil
	}

	for _, reportCard := range reportCards {
		totalGrade += reportCard.Grade
	}

	average := totalGrade / float64(len(reportCards))
	return math.Round(average*10) / 10, nil
}

func (s *StudentServiceImpl) FindTopStudent(subjectName string) (string, error) {
	var topGrades float64
	var topStudent int
	reportCards, err := s.Repo.FindGradesBySubject(subjectName)

	if len(reportCards) <= 0 {
		return "", errors.New("no top student found")
	}

	topGrades = reportCards[0].Grade
	topStudent = reportCards[0].StudentId

	for _, reportCard := range reportCards {
		if reportCard.Grade > topGrades {
			topStudent = reportCard.StudentId
		}
	}

	student, err := s.Repo.FindStudentById(topStudent)
	if err != nil {
		return "", errors.New("no top student found")
	}

	return student.Name, nil
}

func (s *StudentServiceImpl) FilterStudents(id int) (string, error) {
	student, err := s.Repo.FindStudentById(id)
	if err != nil {
		return "", errors.New("no student found")
	}
	return student.Name, nil
}
