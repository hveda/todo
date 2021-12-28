package types

type Response struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type Result struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}


type CreateActivity struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

type CreateToDo struct {
	ActivityGroupId string `json:"activity_group_id"`
	Title string `json:"title"`
	IsActive bool `json:"is_active"`
	Priority string `json:"priority"`
}