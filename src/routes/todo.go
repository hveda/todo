package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hveda/todo/src/types"
	"github.com/hveda/todo/src/utils"
)

func (app *AppBase) TodoItemsGet(w http.ResponseWriter, r *http.Request) {
	todos := make([]types.ToDo, 0)
	app.DB.Table("activities").Find(&todos)
	utils.JsonResponse(w, types.Result{Data: &todos})
}

func (app *AppBase) TodoItemsPost(w http.ResponseWriter, r *http.Request) {
	data := types.CreateToDo{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.JsonResponse(w, types.Response{Message: "Could not parse body data"})
		return
	}
	
	// Create todo
	new_todo := types.ToDo{
		ActivityGroupId: data.ActivityGroupId,
		Title: data.Title,
		IsActive: data.IsActive,
		Priority: data.Priority,
	}
	app.DB.Create(&new_todo)
	utils.JsonResponse(w, new_todo)
}

func (app *AppBase) DeleteToDoById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	todo := types.ToDo{}
	app.DB.Delete(&todo, "id = ?", id)
	if todo.ID == 0 {
		w.WriteHeader(404)
		utils.JsonResponse(w, types.Response{Message: "Could not find activity"})
		return
	}
	utils.JsonResponse(w, &todo)

}

func (app *AppBase) PatchToDoById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	todo := types.ToDo{}
	app.DB.Table("to_dos").Find(&todo, "id = ?", id)
	if todo.ID == 0 {
		w.WriteHeader(404)
		utils.JsonResponse(w, types.Response{Message: "Could not find activity"})
		return
	}
	app.DB.Save(&todo)
	utils.JsonResponse(w, &todo)
}

func (app *AppBase) ShowToDoById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	todo := types.ToDo{}
	app.DB.Table("to_dos").Find(&todo, "id = ?", id)
	if todo.ID == 0 {
		w.WriteHeader(404)
		utils.JsonResponse(w, types.Response{Message: "Could not find activity"})
		return
	}
	utils.JsonResponse(w, &todo)
}
