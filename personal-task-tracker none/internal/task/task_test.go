package task_test

import (
	"personal-task-tracker/internal/task"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestService(t *testing.T) {

	t.Run("add task", func(t *testing.T) {
		service := task.List{}
		assert.Equal(t, 0, len(service.Tasks()))

		someTask := "some task"
		service.Add(someTask)
		service.Add("some")
		service.Add("task")

		assert.Equal(t, 3, len(service.Tasks()))
	})

	t.Run("rename task", func(t *testing.T) {
		service := task.List{}

		service.Add("some description")
		service.Add("bleh")

		tasks := service.Tasks()
		assert.Equal(t, "some description", tasks[0].Description)

		id := tasks[0].ID
		service.Rename(id, "diff desc")

		tasks = service.Tasks()
		assert.Equal(t, "diff desc", tasks[0].Description)
		assert.Equal(t, "bleh", tasks[1].Description)
	})

	t.Run("delete task", func(t *testing.T) {
		service := task.List{}

		service.Add("some1")
		service.Add("some2")
		service.Add("some3")
		service.Add("some4")

		tasks := service.Tasks()
		assert.Equal(t, 4, len(tasks))

		id := tasks[len(tasks)-1].ID

		service.Delete(id)

		tasks = service.Tasks()
		assert.Equal(t, 3, len(tasks))

		assert.Equal(t, "some1", tasks[0].Description)
		assert.Equal(t, "some2", tasks[1].Description)
		assert.Equal(t, "some3", tasks[2].Description)
	})

	t.Run("toggle completion", func(t *testing.T) {
		service := task.List{}

		assert.Equal(t, 0, len(service.Tasks()))

		service.Add("some1")
		service.Add("some2")
		service.Add("some3")

		tasks := service.Tasks()
		assert.Equal(t, 3, len(tasks))

		id := tasks[0].ID
		assert.False(t, tasks[0].Complete)

		service.ToggleDone(id)
		tasks = service.Tasks()
		assert.True(t, tasks[0].Complete)

	})

	t.Run("reoder", func(t *testing.T) {
		service := task.List{}

		service.Add("some1")
		service.Add("some2")
		service.Add("some3")

		tasks := service.Tasks()
		assert.Equal(t, 3, len(tasks))
		assert.Equal(t, "some1", tasks[0].Description)

		//reordering descending
		service.ReOrder([]string{
			tasks[2].ID.String(),
			tasks[1].ID.String(),
			tasks[0].ID.String(),
		})

		tasks = service.Tasks()
		assert.Equal(t, 3, len(tasks))

		assert.Equal(t, "some3", tasks[0].Description)
		assert.Equal(t, "some2", tasks[1].Description)
		assert.Equal(t, "some1", tasks[2].Description)
	})

	t.Run("search by description", func(t *testing.T) {
		service := task.List{}

		service.Add("some1")
		service.Add("some2")
		service.Add("some3")

		tasks := service.Search("som")
		assert.Equal(t, 3, len(tasks))

		assert.Equal(t, "some1", tasks[0].Description)
		assert.Equal(t, "some2", tasks[1].Description)
		assert.Equal(t, "some3", tasks[2].Description)
	})

	t.Run("empty list", func(t *testing.T) {
		service := task.List{}

		service.Add("some1")
		service.Add("some2")
		service.Add("some3")

		tasks := service.Tasks()
		assert.Equal(t, 3, len(tasks))

		service.Empty()

		tasks = service.Tasks()
		assert.Equal(t, 0, len(tasks))
	})
}
