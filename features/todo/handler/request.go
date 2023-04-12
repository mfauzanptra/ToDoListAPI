package handler

import "todolist/features/todo"

type TodoReq struct {
	Title           string `json:"title"`
	ActivityGroupId uint   `json:"activity_group_id"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority"`
}

func ToCore(req *TodoReq) *todo.Core {
	return &todo.Core{
		Title:           req.Title,
		ActivityGroupId: req.ActivityGroupId,
		IsActive:        req.IsActive,
		Priority:        req.Priority,
	}
}
