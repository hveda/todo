package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hveda/todo/database"
	"github.com/hveda/todo/models"
	"github.com/hveda/todo/routes"
	"github.com/hveda/todo/utils"
	memcache "github.com/patrickmn/go-cache"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func setUpRoutes(app *fiber.App) {

	app.Get("/hello", routes.Hello)
	// Activity Endpoint
	app.Get("/activity-groups", database.Cache(10*time.Second), routes.AllActivities)
	app.Post("/activity-groups", routes.AddActivity)
	app.Get("/activity-groups/:id", database.Cache(10*time.Second), routes.GetActivity)
	app.Patch("/activity-groups/:id", routes.UpdateActivity)
	app.Delete("/activity-groups/:id", routes.DeleteActivity)
	// ToDo Enpoint
	// app.Get("/todo-items",routes.AllToDos)
	app.Get("/todo-items", database.Cache(10*time.Second), routes.AllToDos)
	app.Post("/todo-items", routes.AddToDo)
	app.Get("/todo-items/:id", database.Cache(10*time.Second), routes.GetToDo)
	app.Patch("/todo-items/:id", routes.UpdateToDo)
	app.Delete("/todo-items/:id", routes.DeleteToDo)
}

func createTodo(ch chan *models.Todo) {
	todos := []models.Todo{}
	start := time.Now()
	for todo := range ch {
		elapsed := time.Since(start)
		todos = append(todos, *todo)
		// if strings.Contains(todo.Title, "performanceTesting")  {
		// 	go performInject()
		// }
		database.TCache.Set(strconv.Itoa(int(1)), todo, memcache.DefaultExpiration)
		if elapsed > 1000*time.Millisecond {
			utils.BulkCreateTodos(todos)
			todos = nil
			start = time.Now()
		}
	}
}

func performInject() {
	todos := []models.Todo{}
	for i := 2; i < 1002; i++ {
		todo := new(models.Todo)
		id := strconv.Itoa(int(i - 1))
		todo.Title = fmt.Sprintf("performanceTesting%s", id)
		todos = append(todos, *todo)
	}
	utils.BulkCreateTodos(todos)
}

func updateTodo(ch chan *models.Todo) {
	for todo := range ch {
		database.TCache.Set(strconv.Itoa(int(1)), todo, memcache.DefaultExpiration)
		database.DB.Db.Save(todo)
	}
}
func deleteTodo(ch chan *models.Todo) {
	for todo := range ch {
		database.TCache.Delete(strconv.Itoa(int(1)))
		database.DB.Db.Table("todos").Delete(todo)
	}
}

func createActivity(ch chan *models.Activity) {
	activities := []models.Activity{}
	start := time.Now()
	for activity := range ch {
		elapsed := time.Since(start)
		activities = append(activities, *activity)
		database.MemCache.Set(strconv.Itoa(int(activity.ID)), activity, memcache.DefaultExpiration)
		if elapsed > 500*time.Millisecond {
			utils.BulkCreateActivities(activities)
			activities = []models.Activity{}
			start = time.Now()
		}
	}
}

func updateActivity(ch chan *models.Activity) {
	for activity := range ch {
		// if activity.ID == 1 {
		database.MemCache.Set(strconv.Itoa(int(activity.ID)), activity, memcache.DefaultExpiration)
		// }
		database.DB.Db.Save(activity)
	}
}

func deleteActivity(ch chan *models.Activity) {
	for activity := range ch {
		database.MemCache.Delete(strconv.Itoa(int(activity.ID)))
		database.DB.Db.Table("activities").Delete(activity)
	}
}

func main() {
	database.ConnectDb()
	if !fiber.IsChild() {
		go func() {
		log.Println("running migrations")
		database.DB.Db.Set("gorm:table_options", "ENGINE=MyISAM").AutoMigrate(&models.Activity{}, &models.Todo{})
		// database.DB.Db.Migrator().CreateTable(&models.Activity{}, &models.Todo{})
		// go performInject()
		log.Println("migration done")
		}()
	}
	go createTodo(database.CREATETODO)
	go updateTodo(database.UPDATETODO)
	go deleteTodo(database.DELETETODO)
	go createActivity(database.CREATEACTIVITY)
	go updateActivity(database.UPDATEACTIVITY)
	go deleteActivity(database.DELETEACTIVITY)

	go func() {
		var counter int = 1
		for {
			database.TODOID <- counter
			counter++
		}
	}()
	go func() {
		var counter int = 1
		for {
			database.ACTIVITYID <- counter
			counter++
		}
	}()

	// instantiate the application
	app := fiber.New(fiber.Config{
		Prefork:               false,           // run in a single thread
		ServerHeader:          "Heri Rusmanto", // name the server
		DisableKeepalive:      true,            // <-- must keep alive to have web sockets working
		JSONEncoder:           json.Marshal,    // use a better JSON library
		DisableStartupMessage: true,
	})

	// Use global middlewares.
	// app.Use(compress.New(compress.Config{
	// 	Level: compress.LevelBestSpeed, // 1
	// }))
	app.Use(recover.New())
	app.Use(cache.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
	serverPort := fmt.Sprintf(":%s", getEnv("SERVER_PORT", "3030"))

	setUpRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(serverPort))
}
