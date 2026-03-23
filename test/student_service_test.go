package test

import (
	"errors"
	"test/model"
	"test/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
test case:
- add grade to student (name cannot be empty, grade must be <= 100 && >= 0, subject != "")
- calculate average of student (if student has no grade, return 0.0)
- get top student in each subject (if no top student return error)
- filter students (return error if non exist)
*/

var mockRepository = NewMockStudentRepository()
var studentService service.StudentService = &service.StudentServiceImpl{
	Repo: mockRepository,
}

type AddStudentGradeTestCases struct {
	description    string
	student        model.Student
	reportCard     model.ReportCard
	expectedOutput string
	err            error
}

func TestAddStudentGrade(t *testing.T) {
	// Arrange
	testCases := []AddStudentGradeTestCases{
		{
			description: "valid add student grade",
			student: model.Student{
				Name: "Justin",
			},
			reportCard: model.ReportCard{
				Subject:   "History",
				Grade:     89.2,
				StudentId: 1,
			},
			expectedOutput: "Add Success",
			err:            nil,
		},
		{
			description: "invalid add student empty name",
			student: model.Student{
				Name: "",
			},
			reportCard: model.ReportCard{
				Subject:   "History",
				Grade:     89.4,
				StudentId: 1,
			},
			expectedOutput: "",
			err:            errors.New("empty student name"),
		},
		{
			description: "invalid add student larger grade",
			student: model.Student{
				Name: "Justin",
			},
			reportCard: model.ReportCard{
				Subject:   "History",
				Grade:     100.6,
				StudentId: 1,
			},
			expectedOutput: "",
			err:            errors.New("invalid grade"),
		},
		{
			description: "invalid add student smaller grade",
			student: model.Student{
				Name: "Justin",
			},
			reportCard: model.ReportCard{
				Subject:   "History",
				Grade:     -6.5,
				StudentId: 1,
			},
			expectedOutput: "",
			err:            errors.New("invalid grade"),
		},
		{
			description: "invalid add student empty subject",
			student: model.Student{
				Name: "Justin",
			},
			reportCard: model.ReportCard{
				Subject:   "",
				Grade:     89.6,
				StudentId: 1,
			},
			expectedOutput: "",
			err:            errors.New("subject is empty"),
		},
	}

	// Act, Assert
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result, err := studentService.AddStudentGrade(testCase.student, testCase.reportCard)
			assert.Equal(t, testCase.err, err)
			assert.Equal(t, testCase.expectedOutput, result)
		})
	}
}

type CalculateAverageStudentGradeTestCases struct {
	description    string
	studentId      int
	reportCards    []model.ReportCard
	expectedOutput float64
	err            error
}

func TestCalculateAverageGrade(t *testing.T) {
	// Arrange
	mockRepository.AddStudent(model.Student{Id: 1, Name: "Justin"})
	mockRepository.AddGrade(model.ReportCard{
		Id:        2,
		Subject:   "English",
		Grade:     56.7,
		StudentId: 1,
	})
	mockRepository.AddGrade(model.ReportCard{
		Id:        3,
		Subject:   "Math",
		Grade:     22.3,
		StudentId: 1,
	})
	mockRepository.AddGrade(model.ReportCard{
		Id:        4,
		Subject:   "Science",
		Grade:     75.4,
		StudentId: 1,
	})

	testCases := []CalculateAverageStudentGradeTestCases{
		{
			description:    "valid calculate average grade",
			studentId:      1,
			expectedOutput: 60.9,
			err:            nil,
		},
		{
			description:    "invalid empty grade",
			studentId:      2,
			expectedOutput: 0.0,
			err:            nil,
		},
	}

	// Act, Assert
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result, err := studentService.CalculateStudentAverage(testCase.studentId)
			assert.Equal(t, testCase.expectedOutput, result)
			assert.Equal(t, testCase.err, err)
		})
	}
}

type FindTopStudentTestCases struct {
	description    string
	subjectName    string
	expectedOutput string
	err            error
}

func TestFindTopStudent(t *testing.T) {
	// Arrange
	mockRepository.AddStudent(model.Student{Id: 1, Name: "Justin"})
	mockRepository.AddStudent(model.Student{Id: 2, Name: "Amy"})
	mockRepository.AddGrade(model.ReportCard{
		Id:        1,
		Subject:   "History",
		Grade:     89.2,
		StudentId: 1,
	})
	mockRepository.AddGrade(model.ReportCard{
		Id:        2,
		Subject:   "History",
		Grade:     98.6,
		StudentId: 2,
	})

	testCases := []FindTopStudentTestCases{
		{
			description:    "valid find top student",
			subjectName:    "History",
			expectedOutput: "Amy",
			err:            nil,
		},
		{
			description:    "invalid no top student",
			subjectName:    "Chinese",
			expectedOutput: "",
			err:            errors.New("no top student found"),
		},
	}

	// Act, Assert
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result, err := studentService.FindTopStudent(testCase.subjectName)
			assert.Equal(t, testCase.expectedOutput, result)
			assert.Equal(t, testCase.err, err)
		})
	}
}

type FilterStudentsTestCases struct {
	description    string
	studentId      int
	expectedOutput string
	err            error
}

func TestFilterStudents(t *testing.T) {
	// Arrange
	mockRepository.AddStudent(model.Student{Id: 1, Name: "Justin"})
	mockRepository.AddStudent(model.Student{Id: 2, Name: "Amy"})

	testCases := []FilterStudentsTestCases{
		{
			description:    "valid filter students",
			studentId:      2,
			expectedOutput: "Amy",
			err:            nil,
		},
		{
			description:    "invalid filter students",
			studentId:      3,
			expectedOutput: "",
			err:            errors.New("no student found"),
		},
	}

	// Act, Assert
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			result, err := studentService.FilterStudents(testCase.studentId)
			assert.Equal(t, testCase.expectedOutput, result)
			assert.Equal(t, testCase.err, err)
		})
	}
}
