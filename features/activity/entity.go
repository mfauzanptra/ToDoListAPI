package activity

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ActivityId uint      `json:"id"`
	Title      string    `json:"title"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type ActivityHandler interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetOne() echo.HandlerFunc
}

type ActivityService interface {
	Create(newAct Core) (Core, error)
	Update(id uint, updAct Core) (Core, error)
	Delete(id uint) error
	GetAll() ([]Core, error)
	GetOne(id uint) (Core, error)
}

type ActivityData interface {
	Create(newAct Core) (Core, error)
	Update(updAct Core) (Core, error)
	Delete(id uint) error
	GetAll() ([]Core, error)
	GetOne(id uint) (Core, error)
}
