package entity

type User struct {
	Email    string `form:"email"`
	Name     string `form:"name"`
	Password string `form:"password"`
	Tasks    []Task
}

type Task struct {
	Completed   bool   `form:"completed"`
	Title       string `form:"title"`
	Description string `form:"description"`
}
