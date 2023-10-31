package models

type User struct {
	Name     string `json:"name" form:"name" query:"name"`
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"pass" form:"pass" query:"pass"`
	Tasks    []Task `json:"task" form:"task"`
}

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
