package handlers

import (
	"database/sql"
	"github.com/labstack/echo"
	"go-echo-vue/models"
	"net/http"
	"strconv"
)

type H map[string]interface{}

func GetTasks(db *sql.DB) echo.HandlerFunc  {

	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

func PutTasks(db *sql.DB) echo.HandlerFunc  {
	return func(c echo.Context) error {
		// New Task
		var task models.Task

		// Map incoming JSON
		c.Bind(&task)

		// add task to model
		id, err:= models.PutTask(db, task.Name)

		if err == nil {
			return  c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else {
			return err
		}
	}
}

func DeleteTasks(db *sql.DB) echo.HandlerFunc  {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		// Delete task using models
		_, err := models.DeleteTask(db, id)

		if err != nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return  err
		}
	}
}

