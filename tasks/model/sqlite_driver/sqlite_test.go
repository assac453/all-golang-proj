package model_sqlite_test

import (
	"testing"

	"github.com/assac453/tasks/model"
	model_sqlite "github.com/assac453/tasks/model/sqlite_driver"
	"github.com/stretchr/testify/assert"
)

// todo test
func TestSQLLITE(t *testing.T) {
	t.Run("Saving", func(t *testing.T) {
		db := model_sqlite.New()
		assert.Equal(t, 0, db.Count())
		db.Save(model.Task_sqlite{
			ID:        0,
			Title:     "Some",
			Completed: "false",
		})
	})

}
