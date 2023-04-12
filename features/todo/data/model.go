package data

import (
	"time"
	activity "todolist/features/activity/data"
	"todolist/features/todo"
)

type Todo struct {
	TodoId          uint `gorm:"primaryKey"`
	ActivityGroupId uint
	Activity        activity.Activity `gorm:"foreignKey:ActivityGroupId"`
	Title           string
	Priority        string
	IsActive        bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func CoreToData(core todo.Core) Todo {
	return Todo{
		TodoId:          core.TodoId,
		Title:           core.Title,
		ActivityGroupId: core.ActivityGroupId,
		IsActive:        core.IsActive,
		Priority:        core.Priority,
		CreatedAt:       core.CreatedAt,
		UpdatedAt:       core.CreatedAt,
	}
}

func ToCore(data Todo) todo.Core {
	return todo.Core{
		TodoId:          data.TodoId,
		Title:           data.Title,
		ActivityGroupId: data.ActivityGroupId,
		IsActive:        data.IsActive,
		Priority:        data.Priority,
		UpdatedAt:       data.UpdatedAt,
		CreatedAt:       data.CreatedAt,
	}
}
