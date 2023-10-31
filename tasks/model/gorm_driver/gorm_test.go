package model_gorm_test

import (
	"testing"

	"github.com/assac453/tasks/model"
	model_gorm "github.com/assac453/tasks/model/gorm_driver"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func TestModel(t *testing.T) {

	comparer := cmp.Comparer(func(x, y model.Task) bool {
		return x.Title == y.Title && x.Completed == y.Completed
	})

	t.Run("saving", func(t *testing.T) {
		db := model_gorm.New()
		task := model.Task{Title: "title", Completed: "false"}
		db.Save(task)

		count, _ := db.GetAll()
		if len(count) != 1 {
			t.Fail()
		}

		res1, _ := db.GetAll()

		res := res1[0]

		t.Log(task)
		t.Log(res)
		if diff := cmp.Diff(task, res, comparer); diff != "" {
			t.Log(diff)
			t.Fail()
		}

		db.Save(model.Task{
			Title:     "new title",
			Completed: "true",
		})

		res1, _ = db.GetAll()
		res = res1[1]
		if diff := cmp.Diff(model.Task{
			Title:     "new title",
			Completed: "true",
		}, res, comparer); diff != "" {
			t.Fail()
		}
	})
	t.Run("count of rows", func(t *testing.T) {
		db := model_gorm.New()
		assert.Equal(t, 0, db.Count())
		db.Save(model.Task{})
		assert.Equal(t, 1, db.Count())

		for i := 0; i < 10; i++ {
			db.Save(model.Task{})
		}
		assert.Equal(t, 11, db.Count())

		db.Delete(3)
		assert.Equal(t, 10, db.Count())
	})

	t.Run("delete by id", func(t *testing.T) {
		db := model_gorm.New()

		db.Save(model.Task{})
		db.Save(model.Task{})

		db.Delete(1)
		assert.Equal(t, 1, db.Count())

	})

	t.Run("get all", func(t *testing.T) {
		db := model_gorm.New()

		for i := 0; i < 10; i++ {
			db.Save(model.Task{})
		}

		assert.Equal(t, 10, db.Count())

		tasks, _ := db.GetAll()

		assert.Equal(t, 10, len(tasks))

		count, _ := db.GetAll()
		assert.Equal(t, len(count), db.Count())
	})

	t.Run("get by id", func(t *testing.T) {
		db := model_gorm.New()
		tasks := []model.Task{
			{
				Title:     "some title",
				Completed: "false",
			},
			{
				Title:     "some some some",
				Completed: "true",
			},
		}

		db.Save(tasks[0])
		db.Save(tasks[1])

		one, _ := db.Get(1)
		if diff := cmp.Diff(tasks[0], one, comparer); diff != "" {
			t.Fail()
		}

		one, _ = db.Get(2)
		if diff := cmp.Diff(tasks[1], one, comparer); diff != "" {
			t.Fail()
		}
	})

	t.Run("update", func(t *testing.T) {
		db := model_gorm.New()

		tasks := []model.Task{
			{
				Title:     "some",
				Completed: "true",
			},
		}
		db.Save(tasks[0])

		tasks[0].Completed = "false"
		db.Update(1, tasks[0])

		t1, _ := db.Get(1)

		assert.Equal(t, "false", t1.Completed)
		assert.Equal(t, "some", t1.Title)

		tasks[0].Title = "none"
		db.Update(1, tasks[0])

		t1, _ = db.Get(1)

		assert.Equal(t, "none", t1.Title)
		assert.Equal(t, "false", t1.Completed)

	})
}
