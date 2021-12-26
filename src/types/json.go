package types

type Response struct {
	Message string `json:"message"`
}

type Person struct {
	Name string `json:"name"`
	Age uint `json:"age"`
}

type Result struct {
	Data interface{} `json:"data"`
}

type DbConnection struct {
	DbName string `json:"db_name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreatePost struct {
	Title string `json:"title"`
	Content string `json:"content"`
	Summary string `json:"summary"`
}

type CreateActivity struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

type CreateToDo struct {
	ActivityGroupId int64 `json:"activity_group_id"`
	Title string `json:"title"`
	IsActive bool `json:"is_active"`
	Priority string `json:"priority"`
}