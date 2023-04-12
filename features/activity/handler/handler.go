package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"todolist/features/activity"
	"todolist/helper"

	"github.com/labstack/echo/v4"
)

type activityControl struct {
	srv activity.ActivityService
}

func New(srv activity.ActivityService) activity.ActivityHandler {
	return &activityControl{
		srv: srv,
	}
}

func (ac *activityControl) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		newAct := ActReq{}

		if err := c.Bind(&newAct); err != nil {
			log.Println("error bind input")
			return c.JSON(http.StatusBadRequest, helper.ErrResp("Bad Request", "binding error"))
		}

		res, err := ac.srv.Create(*ToCore(&newAct))
		if err != nil {
			if strings.Contains(err.Error(), "title") {
				return c.JSON(http.StatusBadRequest, helper.ErrResp("Bad Request", "title cannot be null"))
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
func (ac *activityControl) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idCnv, _ := strconv.Atoi(id)
		updAct := ActReq{}
		c.Bind(&updAct)

		res, err := ac.srv.Update(uint(idCnv), *ToCore(&updAct))
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, helper.ErrResp("Not Found", "Activity with ID "+id+" Not Found"))
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
func (ac *activityControl) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idCnv, _ := strconv.Atoi(id)

		err := ac.srv.Delete(uint(idCnv))
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, helper.ErrResp("Not Found", "Activity with ID "+id+" Not Found"))
			} else {
				return c.JSON(http.StatusInternalServerError, helper.ErrResp("Internal Server Error", err.Error()))
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "Success",
			"message": "Success",
			"data":    nil,
		})
	}
}
func (ac *activityControl) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := ac.srv.GetAll()
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
func (ac *activityControl) GetOne() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		idCnv, _ := strconv.Atoi(id)

		res, err := ac.srv.GetOne(uint(idCnv))
		if err != nil {
			return c.JSON(http.StatusNotFound, helper.ErrResp("Not Found", "Activity with ID "+id+" Not Found"))
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"status":  "Success",
			"message": "Success",
			"data":    res,
		})
	}
}
