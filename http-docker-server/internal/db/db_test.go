package db_test

import (
	"log"
	"math/rand"
	"testing"

	"github.com/assac453/http-docker-server/internal/db"
	"github.com/google/go-cmp/cmp"
)

func TestDB(t *testing.T) {
	comparer := cmp.Comparer(func(x, y db.Task) bool {
		return x.ID == y.ID && x.Title == y.Title && x.Completed == y.Completed && x.Description == y.Description
	})

	t.Run("add to db", func(t *testing.T) {
		db.InitDB()
		task := db.Task{
			ID:          1,
			Title:       "Some",
			Description: "Testing",
			Completed:   1,
		}
		db.Add(task)
		taskExpected := db.GetAll()

		if len(taskExpected) < 1 {
			t.Fail()
		}

		if diff := cmp.Diff(task, taskExpected[0], comparer); diff != "" {
			t.Error(diff)
		}

		task = db.Task{
			ID:          2,
			Title:       "somesome",
			Description: "some testing for this",
			Completed:   0,
		}
		db.Add(task)
		taskExpected = db.GetAll()

		if len(taskExpected) < 2 {
			t.Fail()
		}

		if diff := cmp.Diff(task, taskExpected[1], comparer); diff != "" {
			t.Error(diff)
		}

	})
	t.Run("get by id", func(t *testing.T) {
		db.InitDB()
		task := db.Task{
			ID:          1,
			Title:       "Some",
			Description: "Testing",
			Completed:   1,
		}
		db.Add(task)
		taskExpected := db.GetById(1)

		if diff := cmp.Diff(task, taskExpected, comparer); diff != "" {
			t.Error(diff)
		}
	})
	t.Run("update by id", func(t *testing.T) {

		db.InitDB()
		task := db.Task{
			ID:          1,
			Title:       "Some",
			Description: "Testing",
			Completed:   1,
		}
		db.Add(task)
		taskExpected := db.GetById(1)

		if diff := cmp.Diff(task, taskExpected, comparer); diff != "" {
			t.Error(diff)
		}

		db.UpdateById(1, db.Task{
			Title:       "some",
			Description: "testing",
			Completed:   0,
		})

		taskExpected = db.GetById(1)

		if !(taskExpected.Title == "some" && taskExpected.Description == "testing" && taskExpected.Completed == 0) {
			t.Fail()
		}
	})

	t.Run("get all task", func(t *testing.T) {
		db.InitDB()

		var tasks []db.Task

		for i := 0; i < 10; i++ {
			task := db.Task{}
			var titleGen []byte

			for j := 0; j < 50; j++ {
				titleGen = append(titleGen, byte(rand.Int()%20+97))
			}

			var descriptionGen []byte
			for j := 0; j < 100; j++ {
				descriptionGen = append(descriptionGen, byte(rand.Int()%20+97))
			}

			task.Title = string(titleGen)
			task.Description = string(descriptionGen)
			task.ID = int64(i)
			task.Completed = int64(rand.Int() % 2)

			tasks = append(tasks, task)
			log.Println(task)
		}

		for _, v := range tasks {
			db.Add(v)
		}

		getAll := db.GetAll()
		if len(getAll) != 10 {
			t.Fail()
		}

	})
	t.Run("delete by id", func(t *testing.T) {
		db.InitDB()
		task := db.Task{
			ID:          1,
			Title:       "Some",
			Description: "Testing",
			Completed:   1,
		}
		db.Add(task)
		taskExpected := db.GetAll()

		if len(taskExpected) < 1 {
			t.Fail()
		}

		db.DeleteById(1)

		taskExpected = db.GetAll()

		if len(taskExpected) > 0 {
			t.Fail()
		}

	})
}
