package routes

import (
	"database/sql"

	"github.com/hveda/todo/src/types"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type AppBase struct {
	DB *gorm.DB
}

type IAppBase interface {
	NewBaseHandler(conn *sql.DB)
	CreateRoutes(router *mux.Router)
}

func (app *AppBase) NewBaseHandler(conn *gorm.DB) *AppBase {
	app.DB = conn

	conn.AutoMigrate(&types.Post{})
	conn.AutoMigrate(&types.Home{})
	return app
}

// Create routes given a gorilla/mux router instance
func (app *AppBase) CreateRoutes(router *mux.Router) *AppBase {
	router.HandleFunc("/", app.Home).Methods("GET")
	router.HandleFunc("/posts", app.CreatePost).Methods("POST")
	router.HandleFunc("/posts", app.GetPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", app.GetPostFromId).Methods("GET")
	// ActivityGroupsGet - Get All
	router.HandleFunc("/activity-groups", app.ActivityGroupsGet).Methods("GET")
	// ActivityGroupsPost - Create
	router.HandleFunc("/activity-groups", app.ActivityGroupsPost).Methods("POST")
	// DeleteActivityById - Delete
	router.HandleFunc("/activity-groups/{id}", app.DeleteActivityById).Methods("DELETE")
	// PatchActivityById - Update
	router.HandleFunc("/activity-groups/{id}", app.PatchActivityById).Methods("PATCH")
	// ShowActivityById - Get One
	router.HandleFunc("/activity-groups/{id}", app.ShowActivityById).Methods("GET")
	// DeleteToDoById - Delete
	router.HandleFunc("/todo-items/{id}", app.DeleteToDoById).Methods("DELETE")
	// PatchToDoById - Update
	router.HandleFunc("/todo-items/{id}", app.PatchToDoById).Methods("PATCh")
	// ShowToDoById - Get One
	router.HandleFunc("/todo-items/{id}", app.ShowToDoById).Methods("GET")
	// TodoItemsGet - Get All
	router.HandleFunc("/todo-items", app.TodoItemsGet).Methods("GET")
	// TodoItemsPost - Create
	router.HandleFunc("/todo-items", app.TodoItemsPost).Methods("POST")

	return app
}

func New(conn *gorm.DB) *AppBase {
	app := AppBase{}
	return app.NewBaseHandler(conn)
}