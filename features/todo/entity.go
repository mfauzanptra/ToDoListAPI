package todo

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Core struct {
	TodoId          uint      `json:"id"`
	Title           string    `json:"title"`
	ActivityGroupId uint      `json:"activity_group_id"`
	IsActive        bool      `json:"is_active"`
	Priority        string    `json:"priority"`
	UpdatedAt       time.Time `json:"updated_at"`
	CreatedAt       time.Time `json:"created_at"`
}

type TodoHandler interface {
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	GetOne() echo.HandlerFunc
}

type TodoService interface {
	Create(newTodo Core) (Core, error)
	Update(id uint, updTodo Core) (Core, error)
	Delete(id uint) error
	GetAll(groupId uint) ([]Core, error)
	GetOne(id uint) (Core, error)
}

type TodoData interface {
	Create(newTodo Core) (Core, error)
	Update(updTodo Core) (Core, error)
	Delete(id uint) error
	GetAll(groupId uint) ([]Core, error)
	GetOne(id uint) (Core, error)
}
