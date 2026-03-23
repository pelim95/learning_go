package test

import (
	"errors"
	"test/model"
)

type MockStudentRepository struct {
	students map[int]model.Student
	reports  []model.ReportCard
}

func NewMockStudentRepository() *MockStudentRepository {
	return &MockStudentRepository{
		students: make(map[int]model.Student),
		reports:  []model.ReportCard{},
	}
}

func (m *MockStudentRepository) AddStudent(student model.Student) (bool, error) {
	m.students[student.Id] = student
	return true, nil
}

func (m *MockStudentRepository) AddGrade(report model.ReportCard) (bool, error) {
	m.reports = append(m.reports, report)
	return true, nil
}

func (m *MockStudentRepository) FindStudentById(id int) (model.Student, error) {
	student, ok := m.students[id]
	if !ok {
		return model.Student{}, errors.New("student not found")
	}
	return student, nil
}

func (m *MockStudentRepository) FindGradesByStudentId(id int) ([]model.ReportCard, error) {
	var result []model.ReportCard
	for _, r := range m.reports {
		if r.StudentId == id {
			result = append(result, r)
		}
	}
	return result, nil
}

func (m *MockStudentRepository) FindGradesBySubject(subject string) ([]model.ReportCard, error) {
	var result []model.ReportCard
	for _, r := range m.reports {
		if r.Subject == subject {
			result = append(result, r)
		}
	}
	return result, nil
}
