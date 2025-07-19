package repositories

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"go-api/.internal/mocks/mockclients"
	"go-api/src/models"
)

type TaskRepositoryTestSuite struct {
	suite.Suite

	db                 *gorm.DB
	mockPostgresClient *mockclients.MockPostgresClient
	repository         TaskRepository
}

func (s *TaskRepositoryTestSuite) SetupTest() {
	t := s.T()
	s.mockPostgresClient = mockclients.NewMockPostgresClient(t)
	// Mock the GetConnection method to return a test database connection
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}
	s.db = db
	s.mockPostgresClient.On("GetConnection").Return(db).Maybe()

	s.repository = NewTaskRepository(TaskRepositoryParams{
		PostgresClient: s.mockPostgresClient,
	})
}

func TestTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}

func (s *TaskRepositoryTestSuite) TestGetTaskByID() {
	tests := map[string]struct {
		setup         func()
		id            int
		expectedTask  *models.Task
		expectedError error
	}{
		"fail - record does not exists": {
			setup: func() {
				s.db.AutoMigrate(&models.Task{})
			},
			id:            1,
			expectedError: gorm.ErrRecordNotFound,
		},
		"success - record exists": {
			setup: func() {
				s.db.AutoMigrate(&models.Task{})
				task := models.Task{
					ID:          1,
					Title:       "Test Task",
					Description: "This is a test task",
					Priority:    1,
					Status:      "todo",
				}
				s.db.Create(&task)
			},
			id: 1,
			expectedTask: &models.Task{
				ID:          1,
				Title:       "Test Task",
				Description: "This is a test task",
				Priority:    1,
				Status:      "todo",
			},
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			s.SetupTest()
			if tc.setup != nil {
				tc.setup()
			}

			task, err := s.repository.GetTaskByID(tc.id)
			if tc.expectedError != nil {
				s.Error(err)
				s.Equal(tc.expectedError, err)
				return
			}
			s.NoError(err)
			s.Equal(tc.expectedTask, task)
		})
	}
}

func (s *TaskRepositoryTestSuite) TestCreateTask() {
	tests := map[string]struct {
		setup         func()
		task          models.Task
		expectedError error
	}{
		"fail - id already exists": {
			setup: func() {
				s.db.AutoMigrate(&models.Task{})
				task := models.Task{
					ID:          1,
					Title:       "Existing Task",
					Description: "This task already exists",
					Priority:    1,
					Status:      "todo",
				}
				s.db.Create(&task)
			},
			task: models.Task{
				ID:          1,
				Title:       "Existing Task",
				Description: "This task already exists",
				Priority:    1,
				Status:      "todo",
			},
			expectedError: errors.New("UNIQUE constraint failed: tasks.id"),
		},
		"success": {
			setup: func() {
				s.db.AutoMigrate(&models.Task{})
			},
			task: models.Task{
				Title:       "New Task",
				Description: "This is a new task",
				Priority:    2,
				Status:      "todo",
			},
			expectedError: nil,
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			s.SetupTest()
			if tc.setup != nil {
				tc.setup()
			}

			err := s.repository.CreateTask(tc.task)
			if tc.expectedError != nil {
				s.Error(err)
				s.Equal(tc.expectedError.Error(), err.Error())
			}
		})
	}
}
