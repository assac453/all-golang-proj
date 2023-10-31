package handlers

import (
	"net/http"

	"github.com/assac453/echo-learning/app/data"
	"github.com/assac453/echo-learning/app/logic"
	"github.com/assac453/echo-learning/app/models"
	"github.com/labstack/echo"
)

func GetUserByEmail(c echo.Context) error {
	u := struct {
		Email string `json:"email"`
	}{}
	if err := c.Bind(&u); err != nil {
		return err
	}
	answer, ok := data.Users[u.Email]
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, answer)
}
func GetAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, data.Users)
}

func AddUserWithotTasks(c echo.Context) error {
	u := models.User{}
	if err := c.Bind(&u); err != nil {
		return err
	}
	data.Users[u.Email] = &u
	return c.JSON(http.StatusOK, u)
}

func AddTaskForUser(c echo.Context) error {
	u := struct {
		Email string        `json:"email"`
		Task  []models.Task `json:"task"`
	}{}
	if err := c.Bind(&u); err != nil {
		return err
	}
	data.Users[u.Email].Tasks = append(data.Users[u.Email].Tasks, u.Task...)
	return c.JSON(http.StatusOK, data.Users[u.Email])
}

func Login(c echo.Context) error {
	// Получение данных аутентификации (например, email и пароль)
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Проверка аутентификационных данных (в твоем случае email и пароль пользователя)
	// Если аутентификация успешна, генерируй токен
	if email == "user@example.com" && password == "password" {
		token, err := logic.GenerateToken(email)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to generate token")
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}

	// Если аутентификация не удалась, возвращай ошибку
	return echo.ErrUnauthorized
}
