package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"todolist/features/todo"
	"todolist/helper"

	"github.com/labstack/echo/v4"
)

type todoControl struct {
	srv todo.TodoService
}

func New(srv todo.TodoService) todo.TodoHandler {
	return &todoControl{
		srv: srv,
	}
}

func (tc *todoControl) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		newTodo := TodoReq{}

		if err := c.Bind(&newTodo); err != nil {
			log.Println("error bind input")
			return c.JSON(http.StatusBadRequest, helper.ErrResp("Bad Request", "binding error"))
		}

		res, err := tc.srv.Create(*ToCore(&newTodo))
		if err != nil {
			if strings.Contains(err.Error(), "title") {
				return c.JSON(http.StatusBadRequest, helper.ErrResp("Bad Request", "title cannot be null"))
			} else if strings.Contains(err.Error(), "group") {
				return c.JSON(http.StatusBadRequest, helper.ErrResp("Bad Request", "activity_group_id cannot be null"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrResp("Internal Server Error", err.Error()))
			}
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"status":  "Success",
			"message": "Success",
			"data":    res,
		})
	}
}
func (tc *todoControl) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idCnv, _ := strconv.Atoi(id)
		updAct := TodoReq{}
		c.Bind(&updAct)

		res, err := tc.srv.Update(uint(idCnv), *ToCore(&updAct))
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, helper.ErrResp("Not Found", "Todo with ID "+id+" Not Found"))
			} else if strings.Contains(err.Error(), "title") {
				return c.JSON(http.StatusBadRequest, helper.ErrResp("Bad Request", "title cannot be null"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrResp("Internal Server Error", err.Error()))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "Success",
			"message": "Success",
			"data":    res,
		})
	}
}
func (tc *todoControl) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idCnv, _ := strconv.Atoi(id)

		err := tc.srv.Delete(uint(idCnv))
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, helper.ErrResp("Not Found", "Todo with ID "+id+" Not Found"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrResp("Internal Server Error", err.Error()))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "Success",
			"message": "Success",
			"data":    map[string]interface{}{},
		})
	}
}
func (tc *todoControl) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		groupId := c.QueryParam("activity_group_id")
		cnvId, _ := strconv.Atoi(groupId)

		res, err := tc.srv.GetAll(uint(cnvId))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrResp("Internal Server Error", err.Error()))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "Success",
			"message": "Success",
			"data":    res,
		})
	}
}
func (tc *todoControl) GetOne() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idCnv, _ := strconv.Atoi(id)

		res, err := tc.srv.GetOne(uint(idCnv))
		if err != nil {
			return c.JSON(http.StatusNotFound, helper.ErrResp("Not Found", "Todo with ID "+id+" Not Found"))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "Success",
			"message": "Success",
			"data":    res,
		})
	}
}
