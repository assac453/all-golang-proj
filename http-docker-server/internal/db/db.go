package db

import (
	"database/sql"
	"log"
	"os"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	m              sync.Mutex
	driverName     string = "sqlite3"
	dataSourceName string = "store.db"
)

type Task struct {
	ID          int64
	Title       string
	Description string
	Completed   int64
}

var exec = `
		CREATE TABLE task(
			id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
			title TEXT NOT NULL,
			description TEXT,
			completed INTEGER NOT NULL
		);
	`

func InitDB() {
	os.Remove(dataSourceName)
	m.Lock()
	defer m.Unlock()

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	_, err = db.Exec(exec)
	if err != nil {
		m.Unlock()
		panic(err)
	}
}

func GetAll() []Task {
	m.Lock()
	defer m.Unlock()

	var tasks []Task

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM task")

	if err != nil {
		m.Unlock()
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		t := Task{}
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed)
		if err != nil {
			log.Println(err)
			continue
		}
		tasks = append(tasks, t)
	}

	for _, v := range tasks {
		log.Println(v)
	}

	return tasks
}

func Add(t Task) {
	m.Lock()
	defer m.Unlock()
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO task (title, description, completed) values ($1,$2,$3)",
		t.Title, t.Description, t.Completed,
	)
	if err != nil {
		m.Unlock()
		log.Println(err)
	}
	if count, err := result.RowsAffected(); err != nil {
		log.Println(count)
	}
}

func GetById(id int64) Task {
	m.Lock()
	defer m.Unlock()

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	rows := db.QueryRow("SELECT * FROM task WHERE ID = $1", id)

	t := Task{}
	err = rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed)
	if err != nil {
		log.Println(err)
	}
	return t
}

func UpdateById(id int64, t Task) {
	m.Lock()
	defer m.Unlock()

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	result, err := db.Exec("UPDATE task SET title=$1, description=$2, completed=$3 WHERE id=$4", t.Title, t.Description, t.Completed, id)
	if err != nil {
		log.Println(err)
	}
	if count, err := result.RowsAffected(); err != nil {
		log.Println(count)
	}
}

func DeleteById(id int64) {
	m.Lock()
	defer m.Unlock()

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	result, err := db.Exec("DELETE FROM task WHERE id = $1", id)
	if err != nil {
		log.Println(err)
	}
	if count, err := result.RowsAffected(); err != nil {
		log.Println(count)
	}
}
