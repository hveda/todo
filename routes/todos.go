package routes

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hveda/todo/database"
	"github.com/hveda/todo/models"
	"github.com/hveda/todo/utils"
)

var empty struct{}
var emptyList = make([]struct{},0)

//AddToDo
func AddToDo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if len(todo.Title) == 0 {
		return c.Status(400).JSON(&fiber.Map{
			"status":  "Bad Request",
			"message": "title cannot be null",
			"data":    empty,
		})
	}

	if todo.ActivityGroupId == 0 {
		return c.Status(400).JSON(&fiber.Map{
			"status":  "Bad Request",
			"message": "activity_group_id cannot be null",
			"data":    empty,
		})
	}

	if todo.Priority == "" {
		todo.Priority = "very-high"
	}
	todo.IsActive = true
	id := <- database.TODOID
	todo.ID = uint16(id)
	database.CREATETODO <- todo
	return c.Status(201).JSON(&fiber.Map{
		"status":  "Success",
		"message": "Success",
		"data":    utils.FormatToDo(*todo),
	})
}

//AllToDos
func AllToDos(c *fiber.Ctx) error {
	todos := []models.Todo{}
	if activity_group_id := c.Query("activity_group_id"); activity_group_id != "" {
		database.DB.Db.Table("todos").Where("activity_group_id = ?", activity_group_id).Find(&todos)
		if len(todos) == 0 {
			return c.Status(200).JSON(&fiber.Map{
				"status":  "Success",
				"message": "Success",
				"data":    make([]struct{},0),
			})
		}
		return c.Status(200).JSON(&fiber.Map{
			"status":  "Success",
			"message": "Success",
			"data":    utils.FormatToDos(todos),
		})
	} else {
		database.DB.Db.Limit(20).Find(&todos)
		return c.Status(200).JSON(&fiber.Map{
			"status":  "Success",
			"message": "Success",
			"data":    utils.FormatToDos(todos),
		})
	}
}

//ToDo
func GetToDo(c *fiber.Ctx) error {
	todo := []models.Todo{}
	if id := c.Params("id"); id != "" {
		if x, found := database.TCache.Get(id); found {
			m := new(models.Todo)
			m = x.(*models.Todo)
			return c.Status(200).JSON(&fiber.Map{
				"status":  "Success",
				"message": "Success",
				"data":    m,
			})
		} else {
			database.DB.Db.Table("todos").Where("id = ?", id).Find(&todo)
			if len(todo) == 0 {
				return c.Status(404).JSON(&fiber.Map{
					"status":  "Not Found",
					"message": fmt.Sprintf("Todo with ID %s Not Found", id),
					"data":    emptyList,
				})
			}
		}
		return c.Status(200).JSON(&fiber.Map{
			"status":  "Success",
			"message": "Success",
			"data":    utils.FormatToDos(todo)[0],
		})
	}
	return c.Status(400).JSON(&fiber.Map{
		"status":  "Bad Request",
		"message": "",
		"data":    empty,
	})
}

//Update
func UpdateToDo(c *fiber.Ctx) error {
	data := new(models.Todo)
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"status":  "Bad Request",
			"message": "",
			"data":    empty,
		})
	}

	todo := []models.Todo{}
	if id := c.Params("id"); id != "" {
		if x, found := database.TCache.Get(id); found {
			m := new(models.Todo)
			m = x.(*models.Todo)
			if len(data.Title) > 0 {
				m.Title = data.Title
			}
			if data.ActivityGroupId == 0 {
				m.ActivityGroupId = data.ActivityGroupId
			}
			if len(data.Priority) > 0 {
				m.Priority = data.Priority
			}
			m.IsActive = data.IsActive
			database.UPDATETODO <- m
			return c.Status(200).JSON(&fiber.Map{
				"status":  "Success",
				"message": "Success",
				"data":    m,
			})
		} else {
			database.DB.Db.Table("todos").Where("id = ?", id).Find(&todo)
			if len(todo) <= 0 {
				return c.Status(404).JSON(&fiber.Map{
					"status":  "Not Found",
					"message": fmt.Sprintf("Todo with ID %s Not Found", id),
					"data":    empty,
				})
			}
			if len(data.Title) > 0 {
				todo[0].Title = data.Title
			}
			if data.ActivityGroupId == 0 {
				todo[0].ActivityGroupId = data.ActivityGroupId
			}
			if len(data.Priority) > 0 {
				todo[0].Priority = data.Priority
			}
			todo[0].IsActive = data.IsActive
			database.UPDATETODO <- &todo[0]
			return c.Status(200).JSON(&fiber.Map{
				"status":  "Success",
				"message": "Success",
				"data":    utils.FormatToDos(todo)[0],
			})
		}
	}
	return c.Status(400).JSON(&fiber.Map{
		"status":  "Bad Request",
		"message": "",
		"data":    empty,
	})
}

//Delete
func DeleteToDo(c *fiber.Ctx) error {
	todo := []models.Todo{}
	if id := c.Params("id"); id != "" {
		if x, found := database.TCache.Get(id); found {
			m := new(models.Todo)
			m = x.(*models.Todo)
			idno, _ := strconv.ParseUint(id, 10, 16)
			m.ID = uint16(idno)
			database.DELETETODO <- m
			return c.Status(200).JSON(&fiber.Map{
				"status":  "Success",
				"message": "Success",
				"data":    empty,
			})
		} else {
			database.DB.Db.Table("todos").Where("id = ?", id).Find(&todo)
			if len(todo) <= 0 {
				return c.Status(404).JSON(&fiber.Map{
					"status":  "Not Found",
					"message": fmt.Sprintf("Todo with ID %s Not Found", id),
					"data":    empty,
				})
			}
		}
		database.DELETETODO <- &todo[0]
		// database.DB.Db.Table("todos").Where("id = ?", id).Delete(&todo[0])
		return c.Status(200).JSON(&fiber.Map{
			"status":  "Success",
			"message": "Success",
			"data":    empty,
		})
	}
	return c.Status(400).JSON(&fiber.Map{
		"status":  "Bad Request",
		"message": "",
		"data":    empty,
	})
}
