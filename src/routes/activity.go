package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/hveda/todo/src/types"
	"github.com/hveda/todo/src/utils"
)

func (app *AppBase) ActivityGroupsGet(w http.ResponseWriter, r *http.Request) {
	activities := make([]types.Activity, 0)
	app.DB.Table("activities").Find(&activities)
	utils.JsonResponse(w, types.Result{Status: "Success", Message: "Success", Data: &activities})
}

func (app *AppBase) ActivityGroupsPost(w http.ResponseWriter, r *http.Request) {
	data := types.CreateActivity{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		utils.JsonResponse(w, types.Response{Status: "Failed", Message: "Could not parse body data", Data: "{}"})
		return
	}
	if len(data.Title) == 0 {
		w.WriteHeader(400)
		w.Header().Add("Content-Type", "application/json")
		utils.JsonResponse(w, types.Result{Status: "400", Message: "Failed to add new activity.", Data: data})
		return
	}
	time_now := time.Now()
	// Create activity
	new_activity := types.Activity{
		Title: data.Title,
		Email: data.Email,
		CreatedAt: time_now,
		UpdatedAt: time_now,
	}
	app.DB.Save(&new_activity)
	// INSERT INTO `activities` (`title`,`email`,`created_at`,`updated_at`,`deleted_at`) VALUES ('','','','',NULL)
	// app.DB.Exec("INSERT INTO `activities` (`title`,`email`,`created_at`,`updated_at`) VALUES (?, ?, ?, ?)", new_activity.Title, new_activity.Email, new_activity.CreatedAt, new_activity.UpdatedAt)
	utils.JsonResponse(w, types.Result{Status: "Success", Message: "Success", Data: new_activity})
}

func (app *AppBase) DeleteActivityById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	activity := types.Activity{}
	app.DB.Delete(&activity, "id = ?", id)
	if activity.ID == 0 {
		w.WriteHeader(404)
		w.Header().Set("Content-Type", "application/json")
		utils.JsonResponse(w, types.Response{
			Status: "Not Found",
			Message: fmt.Sprintf("Activity with ID %s Not Found", id),
			Data: "{}",
		})
		return
	}
	utils.JsonResponse(w, types.Result{Status: "Success", Message: "Success", Data: &activity})
}

func (app *AppBase) PatchActivityById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	activity := types.Activity{}
	app.DB.Table("activities").Find(&activity, "id = ?", id)
	if activity.ID == 0 {
		w.WriteHeader(404)
		w.Header().Set("Content-Type", "application/json")
		utils.JsonResponse(w, types.Response{
			Status: "Not Found",
			Message: fmt.Sprintf("Activity with ID %s Not Found", id),
			Data: "{}",
		})
		return
	}
	app.DB.Save(&activity)
	utils.JsonResponse(w, types.Result{Status: "Success", Message: "Success", Data: &activity})
}

func (app *AppBase) ShowActivityById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	activity := types.Activity{}
	app.DB.Table("activities").Find(&activity, "id = ?", id)
	if activity.ID == 0 {
		w.WriteHeader(404)
		w.Header().Set("Content-Type", "application/json")
		utils.JsonResponse(w, types.Response{
			Status: "Not Found",
			Message: fmt.Sprintf("Activity with ID %s Not Found", id),
			Data: "{}",
		})
		return
	}
	utils.JsonResponse(w, types.Result{Status: "Success", Message: "Success", Data: &activity})
}
