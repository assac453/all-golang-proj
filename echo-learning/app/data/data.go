package data

import "github.com/assac453/echo-learning/app/models"

var Users = map[string]*models.User{
	"example@mail.ru": {
		Name:     "John Silver",
		Email:    "example@mail.ru",
		Password: "somepass",
		Tasks: []models.Task{
			{
				Title:       "Frogs",
				Description: "love frogs",
				Status:      "in progress",
			},
			{
				Title:       "Study",
				Description: "some study",
				Status:      "done",
			},
		},
	},
	"another@example.com": {
		Name:     "Alice Smith",
		Email:    "another@example.com",
		Password: "password123",
		Tasks: []models.Task{
			{
				Title:       "Grocery shopping",
				Description: "buy vegetables and fruits",
				Status:      "in process",
			},
			{
				Title:       "Exercise",
				Description: "go for a jog",
				Status:      "done",
			},
			{
				Title:       "Read book",
				Description: "finish reading 'The Great Gatsby'",
				Status:      "deny",
			},
		},
	},
	"test@example.com": {
		Name:     "Bob Johnson",
		Email:    "test@example.com",
		Password: "test123",
		Tasks: []models.Task{
			{
				Title:       "Clean the house",
				Description: "tidy up the living room",
				Status:      "in process",
			},
			{
				Title:       "Write report",
				Description: "prepare project summary",
				Status:      "deny",
			},
		},
	},
	"user1@example.com": {
		Name:     "User 1",
		Email:    "user1@example.com",
		Password: "password1",
		Tasks:    []models.Task{},
	},
	"user2@example.com": {
		Name:     "User 2",
		Email:    "user2@example.com",
		Password: "password2",
		Tasks: []models.Task{
			{
				Title:       "Task 1",
				Description: "task 1 description",
				Status:      "done",
			},
			{
				Title:       "Task 2",
				Description: "task 2 description",
				Status:      "in process",
			},
		},
	},
	"user3@example.com": {
		Name:     "User 3",
		Email:    "user3@example.com",
		Password: "password3",
		Tasks: []models.Task{
			{
				Title:       "Task 1",
				Description: "task 1 description",
				Status:      "deny",
			},
			{
				Title:       "Task 2",
				Description: "task 2 description",
				Status:      "done",
			},
			{
				Title:       "Task 3",
				Description: "task 3 description",
				Status:      "in process",
			},
		},
	},
	"user4@example.com": {
		Name:     "User 4",
		Email:    "user4@example.com",
		Password: "password4",
		Tasks:    []models.Task{},
	},
	"user5@example.com": {
		Name:     "User 5",
		Email:    "user5@example.com",
		Password: "password5",
		Tasks: []models.Task{
			{
				Title:       "Task 1",
				Description: "task 1 description",
				Status:      "deny",
			},
		},
	},
	"user6@example.com": {
		Name:     "User 6",
		Email:    "user6@example.com",
		Password: "password6",
		Tasks: []models.Task{
			{
				Title:       "Task 1",
				Description: "task 1 description",
				Status:      "in process",
			},
			{
				Title:       "Task 2",
				Description: "task 2 description",
				Status:      "done",
			},
			{
				Title:       "Task 3",
				Description: "task 3 description",
				Status:      "deny",
			},
			{
				Title:       "Task 4",
				Description: "task 4 description",
				Status:      "in process",
			},
		},
	},
	"user7@example.com": {
		Name:     "User 7",
		Email:    "user7@example.com",
		Password: "password7",
		Tasks: []models.Task{
			{
				Title:       "Task 1",
				Description: "task 1 description",
				Status:      "done",
			},
			{
				Title:       "Task 2",
				Description: "task 2 description",
				Status:      "in process",
			},
			{
				Title:       "Task 3",
				Description: "task 3 description",
				Status:      "deny",
			},
			{
				Title:       "Task 4",
				Description: "task 4 description",
				Status:      "done",
			},
			{
				Title:       "Task 5",
				Description: "task 5 description",
				Status:      "in process",
			},
		},
	},
	"user8@example.com": {
		Name:     "User 8",
		Email:    "user8@example.com",
		Password: "password8",
		Tasks: []models.Task{
			{
				Title:       "Task 1",
				Description: "task 1 description",
				Status:      "deny",
			},
			{
				Title:       "Task 2",
				Description: "task 2 description",
				Status:      "in process",
			},
			{
				Title:       "Task 3",
				Description: "task 3 description",
				Status:      "done",
			},
			{
				Title:       "Task 4",
				Description: "task 4 description",
				Status:      "in process",
			},
			{
				Title:       "Task 5",
				Description: "task 5 description",
				Status:      "deny",
			},
			{
				Title:       "Task 6",
				Description: "task 6 description",
				Status:      "done",
			},
		},
	},
	"user9@example.com": {
		Name:     "User 9",
		Email:    "user9@example.com",
		Password: "password9",
		Tasks: []models.Task{
			{
				Title:       "Task 1",
				Description: "task 1 description",
				Status:      "done",
			},
			{
				Title:       "Task 2",
				Description: "task 2 description",
				Status:      "deny",
			},
			{
				Title:       "Task 3",
				Description: "task 3 description",
				Status:      "in process",
			},
			{
				Title:       "Task 4",
				Description: "task 4 description",
				Status:      "done",
			},
			{
				Title:       "Task 5",
				Description: "task 5 description",
				Status:      "deny",
			},
			{
				Title:       "Task 6",
				Description: "task 6 description",
				Status:      "in process",
			},
			{
				Title:       "Task 7",
				Description: "task 7 description",
				Status:      "done",
			},
		},
	},
	"user10@example.com": {
		Name:     "User 10",
		Email:    "user10@example.com",
		Password: "password10",
		Tasks: []models.Task{
			{
				Title:       "Task 1",
				Description: "task 1 description",
				Status:      "in process",
			},
			{
				Title:       "Task 2",
				Description: "task 2 description",
				Status:      "deny",
			},
			{
				Title:       "Task 3",
				Description: "task 3 description",
				Status:      "done",
			},
			{
				Title:       "Task 4",
				Description: "task 4 description",
				Status:      "in process",
			},
			{
				Title:       "Task 5",
				Description: "task 5 description",
				Status:      "done",
			},
			{
				Title:       "Task 6",
				Description: "task 6 description",
				Status:      "in process",
			},
			{
				Title:       "Task 7",
				Description: "task 7 description",
				Status:      "deny",
			},
			{
				Title:       "Task 8",
				Description: "task 8 description",
				Status:      "done",
			},
			{
				Title:       "Task 9",
				Description: "task 9 description",
				Status:      "in process",
			},
			{
				Title:       "Task 10",
				Description: "task 10 description",
				Status:      "done",
			},
		},
	},
}
