package handler

import "todolist/features/activity"

type ActReq struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

func ToCore(req *ActReq) *activity.Core {
	return &activity.Core{
		Title: req.Title,
		Email: req.Email,
	}
}
