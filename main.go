package main

import (
	"log"
	"todolist/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	ad "todolist/features/activity/data"
	ah "todolist/features/activity/handler"
	as "todolist/features/activity/services"

	td "todolist/features/todo/data"
	th "todolist/features/todo/handler"
	ts "todolist/features/todo/services"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)
	config.Migrate(db)

	actData := ad.New(db)
	actSrv := as.New(actData)
	actHdl := ah.New(actSrv)

	todoData := td.New(db)
	todoSrv := ts.New(todoData)
	todoHdl := th.New(todoSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "- method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	//Activity
	e.POST("/activity-groups", actHdl.Create())
	e.PATCH("/activity-groups/:id", actHdl.Update())
	e.DELETE("/activity-groups/:id", actHdl.Delete())
	e.GET("/activity-groups", actHdl.GetAll())
	e.GET("/activity-groups/:id", actHdl.GetOne())

	//Todo
	e.POST("/todo-items", todoHdl.Create())
	e.PATCH("/todo-items/:id", todoHdl.Update())
	e.DELETE("/todo-items/:id", todoHdl.Delete())
	e.GET("/todo-items", todoHdl.GetAll())
	e.GET("/todo-items/:id", todoHdl.GetOne())

	if err := e.Start(":3030"); err != nil {
		log.Println(err.Error())
	}
}
