package data

import (
	"time"
	"todolist/features/activity"

	"gorm.io/gorm"
)

type Activity struct {
	ActivityId uint `gorm:"primaryKey"`
	Title      string
	Email      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

func CoreToData(core activity.Core) Activity {
	return Activity{
		ActivityId: core.ActivityId,
		Title:      core.Title,
		Email:      core.Email,
		CreatedAt:  core.CreatedAt,
		UpdatedAt:  core.UpdatedAt,
	}
}

func ToCore(data Activity) activity.Core {
	return activity.Core{
		ActivityId: data.ActivityId,
		Title:      data.Title,
		Email:      data.Email,
		CreatedAt:  data.CreatedAt,
		UpdatedAt:  data.UpdatedAt,
	}
}
