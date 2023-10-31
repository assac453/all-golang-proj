package model_sqlite

import (
	"database/sql"
	"log"
	"os"

	"github.com/assac453/tasks/model"
	_ "github.com/mattn/go-sqlite3"
)

var (
	schema = "sqlite3"
	store  = "store.db"

	exec = `
		CREATE TABLE task(
			id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			title TEXT NOT NULL,
			completed TEXT NOT NULL
		);
	`
)

type DBsqlite struct {
}

func New() model.IModel {
	d := &DBsqlite{}
	d._init()
	return d
}

func open() (*sql.DB, func()) {
	db, err := sql.Open(schema, store)
	if err != nil {
		log.Println(err)
	}
	return db, func() {
		db.Close()
	}
}

func (d DBsqlite) _init() {
	os.Remove(schema)

	db, close := open()
	defer close()

	_, err := db.Exec(exec)
	if err != nil {
		panic(err)
	}
}
func (d DBsqlite) GetAll() ([]model.Task, error) {
	db, close := open()
	defer close()

	rows, err := db.Query("SELECT * FROM task")
	if err != nil {
		return []model.Task{}, err
	}
	defer rows.Close()
	tasks := []model.Task{}
	for rows.Next() {
		t := model.Task_sqlite{}
		err := rows.Scan(&t.ID, &t.Title, &t.Completed)
		if err != nil {
			log.Println(err)
			continue
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}
func (d DBsqlite) Get(id int) (model.Task, error) {
	db, close := open()
	defer close()
	row := db.QueryRow("SELECT * FROM task WHERE ID=$1", id)
	t := model.Task_sqlite{}
	err := row.Scan(&t.ID, &t.Title, &t.Completed)
	if err != nil {
		return nil, err
	}
	return t, nil
}
func (d DBsqlite) Delete(id int) error {
	db, close := open()
	defer close()

	_, err := db.Exec("DELETE FROM task WHERE ID=$1", id)
	if err != nil {
		return err
	}
	return nil
}
func (d DBsqlite) Update(id int, new model.Task) error {
	db, close := open()
	defer close()
	_, err := db.Exec("UPDATE task SET title=$1, completed=$2 WHERE ID=$3", new.Title, new.Completed, id)
	if err != nil {
		return err
	}
	return nil
}
func (d DBsqlite) Save(t model.Task) error {
	db, close := open()
	defer close()
	_, err := db.Exec("INSERT INTO task (title, completed) VALUES($1,$2)", t.Title, t.Completed)
	if err != nil {
		return err
	}
	return nil
}
func (d DBsqlite) Count() int {
	db, close := open()
	defer close()
	res, err := db.Exec("SELECT * FROM task")
	if err != nil {
		log.Println(err)
		return -1
	}
	count, _ := res.RowsAffected()
	return int(count)
}
