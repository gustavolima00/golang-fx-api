package models

import (
	"time"
)

type Task struct {
	ID          uint
	Title       string
	Description string
	Priority    int
	DueDate     *time.Time
	Status      string
}
