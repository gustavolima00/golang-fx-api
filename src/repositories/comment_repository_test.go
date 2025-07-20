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

type CommentRepositoryTestSuite struct {
	suite.Suite

	db                 *gorm.DB
	mockPostgresClient *mockclients.MockPostgresClient
	repository         CommentRepository
}

func (s *CommentRepositoryTestSuite) SetupTest() {
	t := s.T()
	s.mockPostgresClient = mockclients.NewMockPostgresClient(t)
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}
	s.db = db
	s.mockPostgresClient.On("GetConnection").Return(db).Maybe()

	s.repository = NewCommentRepository(CommentRepositoryParams{
		PostgresClient: s.mockPostgresClient,
	})
}

func TestCommentRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CommentRepositoryTestSuite))
}

func (s *CommentRepositoryTestSuite) TestGetTaskComments() {
	tests := map[string]struct {
		setup            func()
		taskId           int
		expectedResponse []models.Comment
	}{
		"no commnets": {
			setup: func() {
				s.db.AutoMigrate(&models.Comment{})
			},
			taskId:           1,
			expectedResponse: []models.Comment{},
		},
		"success - record exists": {
			setup: func() {
				s.db.AutoMigrate(&models.Comment{})
				task := models.Comment{
					ID:      1,
					TaskID:  1,
					Content: "This is a test comment",
					Author:  "Test User",
				}
				s.db.Create(&task)
			},
			taskId: 1,
			expectedResponse: []models.Comment{
				{
					ID:      1,
					TaskID:  1,
					Content: "This is a test comment",
					Author:  "Test User",
				},
			},
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			s.SetupTest()
			if tc.setup != nil {
				tc.setup()
			}

			res := s.repository.GetTaskComments(tc.taskId)
			s.Equal(tc.expectedResponse, res)
		})
	}
}

func (s *CommentRepositoryTestSuite) TestCreateComment() {
	tests := map[string]struct {
		setup         func()
		input         models.Comment
		expectedError error
	}{
		"fail - id already exists": {
			setup: func() {
				s.db.AutoMigrate(&models.Comment{})
				task := models.Comment{
					ID:      1,
					TaskID:  1,
					Content: "Existing Comment",
					Author:  "Test User",
				}
				s.db.Create(&task)
			},
			input: models.Comment{
				ID:      1,
				TaskID:  1,
				Content: "This comment already exists",
				Author:  "Test User",
			},
			expectedError: errors.New("UNIQUE constraint failed: comments.id"),
		},
		"success": {
			setup: func() {
				s.db.AutoMigrate(&models.Comment{})
			},
			input: models.Comment{
				ID:      1,
				TaskID:  1,
				Content: "This is a new comment",
				Author:  "Test User",
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

			err := s.repository.CreateComment(tc.input)
			if tc.expectedError != nil {
				s.Error(err)
				s.Equal(tc.expectedError.Error(), err.Error())
			}
		})
	}
}
