package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/hveda/todo/database"
	"github.com/hveda/todo/models"
)

type ActivityResponse struct {
	ID        int            `json:"id"`
	Title     string         `json:"title"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func FormatActivity(activity models.Activity) ActivityResponse {
	formatter := ActivityResponse{}
	formatter.ID = int(activity.ID)
	formatter.Title = activity.Title
	formatter.Email = activity.Email
	formatter.CreatedAt = activity.CreatedAt
	formatter.UpdatedAt = activity.UpdatedAt

	return formatter
}

func FormatActivities(activities []models.Activity) []ActivityResponse {
	if len(activities) == 0 {
		return []ActivityResponse{}
	}

	var activitiesFormatter []ActivityResponse

	for _, activity := range activities {
		formatter := FormatActivity(activity)
		activitiesFormatter = append(activitiesFormatter, formatter)
	}

	return activitiesFormatter
}

func BulkCreateActivities(rs []models.Activity) error {
	valueStrings := []string{}
	valueArgs := []interface{}{}

	for _, f := range rs {
		valueStrings = append(valueStrings, "(?, ?)")

		valueArgs = append(valueArgs, f.Title)
		valueArgs = append(valueArgs, f.Email)
	}

	smt := `INSERT INTO activities(title,email) VALUES %s ON DUPLICATE KEY UPDATE title=VALUES(title),email=VALUES(email)`

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