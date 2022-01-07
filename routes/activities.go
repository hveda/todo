package routes

import (
	"fmt"	
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hveda/todo/database"
	"github.com/hveda/todo/models"
	"github.com/hveda/todo/utils"
)

//AddActivity
func AddActivity(c *fiber.Ctx) error {
	activity := new(models.Activity)
	if err := c.BodyParser(activity); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if len(activity.Title) == 0 {
		return c.Status(400).JSON(&fiber.Map{
			"status": "Bad Request",
			"message": "title cannot be null",
			"data": empty,
		})
	}

	id := <- database.ACTIVITYID
	activity.ID = uint16(id)
	database.CREATEACTIVITY <- activity
	return c.Status(201).JSON(&fiber.Map{
		"status": "Success",
		"message": "Success",
		"data": utils.FormatActivity(*activity),
	  })
	return nil
}

//AllActivities
func AllActivities(c *fiber.Ctx) error {
	activities := []models.Activity{}
	database.DB.Db.Find(&activities)

	return c.Status(200).JSON(&fiber.Map{
		"status": "Success",
		"message": "Success",
		"data": utils.FormatActivities(activities),
	  })
}

//Activity
func GetActivity(c *fiber.Ctx) error {
	activity := []models.Activity{}
	if id := c.Params("id"); id != "" {
		if x, found := database.MemCache.Get(id); found {
			m := new(models.Activity)
			m = x.(*models.Activity)
			idno, _ := strconv.ParseUint(id, 10, 16)
			m.ID = uint16(idno)
			return c.Status(200).JSON(&fiber.Map{
				"status":  "Success",
				"message": "Success",
				"data":    utils.FormatActivity(*m),
			})
		} else {
			database.DB.Db.Table("activities").Where("id = ?", id).Find(&activity)
			if len(activity) <= 0 {
				return c.Status(404).JSON(&fiber.Map{
					"status": "Not Found",
					"message": fmt.Sprintf("Activity with ID %s Not Found", id),
					"data": empty,
				})
			}
		}
		return c.Status(200).JSON(&fiber.Map{
			"status": "Success",
			"message": "Success",
			"data": utils.FormatActivities(activity)[0],
		})
	}
	return c.Status(400).JSON(&fiber.Map{
		"status": "Bad Request",
		"message": "",
		"data": empty,
	  })
}

//Update
func UpdateActivity(c *fiber.Ctx) error {
	data := new(models.Activity)
	if err := c.BodyParser(data); err != nil {
			return c.Status(400).JSON(&fiber.Map{
				"status": "Bad Request",
				"message": "",
				"data": empty,
			  })
		}
	activity := []models.Activity{}
	if id := c.Params("id"); id != "" {
		if x, found := database.MemCache.Get(id); found {
			m := new(models.Activity)
			m = x.(*models.Activity)
			idno, _ := strconv.ParseUint(id, 10, 16)
			m.ID = uint16(idno)
			if len(data.Title) > 0 {
				m.Title = data.Title
			}
			if len(data.Email) > 0 {
				m.Email = data.Email
			}
			database.UPDATEACTIVITY <- m
			return c.Status(200).JSON(&fiber.Map{
				"status":  "Success",
				"message": "Success",
				"data":    utils.FormatActivity(*m),
			})
		} else {
			database.DB.Db.Table("activities").Where("id = ?", id).Find(&activity)
			if len(activity) <= 0 {
				return c.Status(404).JSON(&fiber.Map{
					"status": "Not Found",
					"message": fmt.Sprintf("Activity with ID %s Not Found", id),
					"data": empty,
					})
			}
			if len(data.Email) > 0 {
				activity[0].Email = data.Email
			}
			if len(data.Title) > 0 {
				activity[0].Title = data.Title
			}
			database.UPDATEACTIVITY <- &activity[0]
			// database.DB.Db.Save(&activity[0])
			return c.Status(200).JSON(&fiber.Map{
				"status": "Success",
				"message": "Success",
				"data": utils.FormatActivities(activity)[0],
			})
		}
	}
	return c.Status(404).JSON(&fiber.Map{
		"status": "Bad Request",
		"message": "",
		"data": empty,
		})
}

//Delete
func DeleteActivity(c *fiber.Ctx) error {
	activity := []models.Activity{}
	if id := c.Params("id"); id != "" {
		if x, found := database.MemCache.Get(id); found {
			m := new(models.Activity)
			m = x.(*models.Activity)
			idno, _ := strconv.ParseUint(id, 10, 16)
			m.ID = uint16(idno)
			database.DELETEACTIVITY <- m
			return c.Status(200).JSON(&fiber.Map{
				"status":  "Success",
				"message": "Success",
				"data":    empty,
			})
		} else {
			database.DB.Db.Table("activities").Where("id = ?", id).Find(&activity)
			if len(activity) == 0 {
				return c.Status(404).JSON(&fiber.Map{
					"status": "Not Found",
					"message": fmt.Sprintf("Activity with ID %s Not Found", id),
					"data": empty,
					})
				}
		}
		database.DELETEACTIVITY <- &activity[0]
		// go database.DB.Db.Table("activities").Where("id = ?", id).Delete(&activity[0])
		return c.Status(200).JSON(&fiber.Map{
			"status": "Success",
			"message": "Success",
			"data": empty,
		})
	}
	return c.Status(404).JSON(&fiber.Map{
		"status": "Bad Request",
		"message": "",
		"data": empty,
		})
}
