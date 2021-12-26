package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hveda/todo/src/types"
	"github.com/hveda/todo/src/utils"
)

func (app *AppBase) ActivityGroupsGet(w http.ResponseWriter, r *http.Request) {
	activities := make([]types.Activity, 0)
	app.DB.Table("activities").Find(&activities)
	utils.JsonResponse(w, types.Result{Data: &activities})
}

func (app *AppBase) ActivityGroupsPost(w http.ResponseWriter, r *http.Request) {
	data := types.CreateActivity{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.JsonResponse(w, types.Response{Message: "Could not parse body data"})
		return
	}
	
	// Create activity
	new_activity := types.Activity{
		Title: data.Title,
		Email: data.Email,
	}
	app.DB.Create(&new_activity)
	utils.JsonResponse(w, new_activity)
}

func (app *AppBase) DeleteActivityById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	activity := types.Activity{}
	app.DB.Delete(&activity, "id = ?", id)
	if activity.ID == 0 {
		w.WriteHeader(404)
		utils.JsonResponse(w, types.Response{Message: "Could not find activity"})
		return
	}
	utils.JsonResponse(w, &activity)

}

func (app *AppBase) PatchActivityById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	activity := types.Activity{}
	app.DB.Table("activities").Find(&activity, "id = ?", id)
	if activity.ID == 0 {
		w.WriteHeader(404)
		utils.JsonResponse(w, types.Response{Message: "Could not find activity"})
		return
	}
	app.DB.Save(&activity)
	utils.JsonResponse(w, &activity)
}

func (app *AppBase) ShowActivityById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	activity := types.Activity{}
	app.DB.Table("activities").Find(&activity, "id = ?", id)
	if activity.ID == 0 {
		w.WriteHeader(404)
		utils.JsonResponse(w, types.Response{Message: "Could not find activity"})
		return
	}
	utils.JsonResponse(w, &activity)
}
