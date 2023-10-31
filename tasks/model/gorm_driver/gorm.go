package model_gorm

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/assac453/tasks/model"
	zerolog "github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	db *gorm.DB
}

func New() model.IModel {
	d := &DB{}
	d._init()
	return d
}

func (d DB) _init() {
	os.Remove("base.db")
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Silent,
			SlowThreshold:             time.Second,
		},
	)

	var err error
	d.db, err = gorm.Open(sqlite.Open("base.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connection db")
	}

	d.db.AutoMigrate(&model.Task{})
}
func (d DB) Save(t model.Task) error {
	zerolog.Info().Msg(fmt.Sprintf("saving model.Task: %v", t.String()))
	res := d.db.Save(&t)
	if res.Error != nil {
		return res.Error
	}
	zerolog.Info().Msg(fmt.Sprintf("rows affected after saving : %v", res.RowsAffected))
	return nil
}
func (d DB) GetAll() ([]model.Task, error) {
	var tasks []model.Task
	res := d.db.Find(&tasks)

	if res.Error != nil {
		zerolog.Error().Msg(fmt.Sprintf("error occur while getAll:  %v", res.Error.Error()))
		return []model.Task{}, res.Error
	}
	zerolog.Info().Msg(fmt.Sprintf("get all tasks with %v", tasks))
	return tasks, nil
}
func (d DB) Get(id int) (model.Task, error) {
	var t model.Task
	res := d.db.First(&t, id)
	if res.Error != nil {
		zerolog.Error().Msg(fmt.Sprintf("error occur while get by id: %v", res.Error.Error()))
		return model.Task{}, res.Error
	}
	zerolog.Info().Msg(fmt.Sprintf("get task by id %v : %v", id, t))
	return t, nil
}
func (d DB) Count() int {
	var tasks []model.Task
	d.db.Find(&tasks)
	return len(tasks)
}
func (d DB) Delete(id int) error {
	res := d.db.Delete(&model.Task{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
func (d DB) Update(id int, new model.Task) error {
	old, err := d.Get(id)
	if err != nil {
		return err
	}
	d.db.Model(&old).Updates(new)
	return nil
}
