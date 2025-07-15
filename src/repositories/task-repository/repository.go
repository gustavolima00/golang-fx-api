package taskrepository

import (
	"go-api/src/clients/postgres"
	models "go-api/src/models/task"

	"go.uber.org/fx"
)

type TaskRepository interface {
	GetTaskByID(id int) (*models.Task, error)
	CreateTask(task models.Task) error
}

type Params struct {
	fx.In
	PostgresClient postgres.Client
}

type repository struct {
	postgresClient postgres.Client
}

func New(p Params) TaskRepository {
	return &repository{
		postgresClient: p.PostgresClient,
	}
}

func (r *repository) GetTaskByID(id int) (*models.Task, error) {
	db := r.postgresClient.GetConnection()
	var task models.Task
	result := db.Model(&models.Task{}).Where("id = ?", id).First(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (r *repository) CreateTask(task models.Task) error {
	db := r.postgresClient.GetConnection()
	result := db.Create(&task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
