package models

type Comment struct {
	ID      uint
	TaskID  uint
	Content string
	Author  string
}
