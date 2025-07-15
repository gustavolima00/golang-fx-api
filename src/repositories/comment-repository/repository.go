package commnetrepository

import (
	"go-api/src/clients/postgres"
	models "go-api/src/models/comment"

	"go.uber.org/fx"
)

type CommentRepository interface {
	GetTaskComments(id int) []models.Comment
	CreateComment(comment models.Comment) error
}

type Params struct {
	fx.In
	PostgresClient postgres.Client
}

type repository struct {
	postgresClient postgres.Client
}

func New(p Params) CommentRepository {
	return &repository{
		postgresClient: p.PostgresClient,
	}
}

func (r *repository) GetTaskComments(taskId int) []models.Comment {
	db := r.postgresClient.GetConnection()
	var comments []models.Comment
	db.Model(&models.Comment{}).Where("task_id = ?", taskId).Find(&comments)
	return comments
}

func (r *repository) CreateComment(comment models.Comment) error {
	db := r.postgresClient.GetConnection()
	result := db.Create(&comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
