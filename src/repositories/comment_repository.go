package repositories

import (
	"go-api/src/clients"
	"go-api/src/models"

	"go.uber.org/fx"
)

type CommentRepository interface {
	GetTaskComments(id int) []models.Comment
	CreateComment(comment models.Comment) error
}

type CommentRepositoryParams struct {
	fx.In
	PostgresClient clients.PostgresClient
}

type commentRepository struct {
	postgresClient clients.PostgresClient
}

func NewCommentRepository(p CommentRepositoryParams) CommentRepository {
	return &commentRepository{
		postgresClient: p.PostgresClient,
	}
}

func (r *commentRepository) GetTaskComments(taskId int) []models.Comment {
	db := r.postgresClient.GetConnection()
	var comments []models.Comment
	db.Model(&models.Comment{}).Where("task_id = ?", taskId).Find(&comments)
	return comments
}

func (r *commentRepository) CreateComment(comment models.Comment) error {
	db := r.postgresClient.GetConnection()
	result := db.Create(&comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
