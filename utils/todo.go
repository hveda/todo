package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/hveda/todo/database"
	"github.com/hveda/todo/models"
)

type ToDoResponse struct {
	ID              int            `json:"id"`
	ActivityGroupId int            `json:"activity_group_id"`
	Title           string         `json:"title"`
	IsActive        bool           `json:"is_active"`
	Priority        string         `json:"priority"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}

func FormatToDo(todo models.Todo) ToDoResponse {
	formatter := ToDoResponse{}
	formatter.ID = int(todo.ID)
	formatter.ActivityGroupId = int(todo.ActivityGroupId)
	formatter.Title = todo.Title
	formatter.Priority = todo.Priority
	formatter.IsActive = todo.IsActive
	formatter.CreatedAt = todo.CreatedAt
	formatter.UpdatedAt = todo.UpdatedAt

	return formatter
}

func FormatToDos(todos []models.Todo) []ToDoResponse {
	if len(todos) == 0 {
		return []ToDoResponse{}
	}

	var todosFormatter []ToDoResponse

	for _, todo := range todos {
		formatter := FormatToDo(todo)
		todosFormatter = append(todosFormatter, formatter)
	}

	return todosFormatter
}

func BulkCreateTodos(rs []models.Todo) error {
	valueStrings := []string{}
	valueArgs := []interface{}{}

	for _, f := range rs {
		valueStrings = append(valueStrings, "(?, ?, ?, ?)")

		valueArgs = append(valueArgs, f.ActivityGroupId)
		valueArgs = append(valueArgs, f.Title)
		valueArgs = append(valueArgs, true)
		valueArgs = append(valueArgs, "very-high")
	}

	// smt := `INSERT INTO todos(activity_group_id,title,is_active,priority) VALUES %s ON DUPLICATE KEY UPDATE activity_group_id=VALUES(activity_group_id),title=VALUES(title),is_active=VALUES(is_active),priority=VALUES(priority)`
	smt := `INSERT INTO todos(activity_group_id,title,is_active,priority) VALUES %s`

	smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))

	go func() error {
		tx := database.DB.Db.Begin()
		if err := tx.Exec(smt, valueArgs...).Error; err != nil {
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	}()
	return nil
}