package repositories

import (
	"go-api/src/clients"
	"go-api/src/models"

	"go.uber.org/fx"
)

type TaskRepository interface {
	GetTaskByID(id int) (*models.Task, error)
	CreateTask(task models.Task) error
}

type TaskRepositoryParams struct {
	fx.In
	PostgresClient clients.PostgresClient
}

type taskRepository struct {
	postgresClient clients.PostgresClient
}

func NewTaskRepository(p TaskRepositoryParams) TaskRepository {
	return &taskRepository{
		postgresClient: p.PostgresClient,
	}
}

func (r *taskRepository) GetTaskByID(id int) (*models.Task, error) {
	db := r.postgresClient.GetConnection()
	var task models.Task
	result := db.Model(&models.Task{}).Where("id = ?", id).First(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (r *taskRepository) CreateTask(task models.Task) error {
	db := r.postgresClient.GetConnection()
	result := db.Create(&task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
