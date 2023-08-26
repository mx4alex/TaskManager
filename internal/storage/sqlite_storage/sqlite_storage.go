package sqlite_storage

import (
	"TaskManager/internal/entity"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const tasksDB = "data/tasks.db"

type SQLiteStorage struct {
	DB *sql.DB
}

func New() (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", tasksDB)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, text TEXT, done INTEGER)")
	if err != nil {
		return nil, err
	}

	return &SQLiteStorage{DB: db}, nil
}

func (s *SQLiteStorage) AddTask(newText string) error {
	_, err := s.DB.Exec("INSERT INTO tasks (text, done) VALUES (?, ?)", newText, 0)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLiteStorage) GetTasks() ([]entity.Task, error) {
	rows, err := s.DB.Query("SELECT id, text, done FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []entity.Task
	for rows.Next() {
		var id int
		var text string
		var done int

		err := rows.Scan(&id, &text, &done)
		if err != nil {
			return nil, err
		}

		if done == 0 {
			tasks = append(tasks, entity.NewTask(id, text, false))
		} else {
			tasks = append(tasks, entity.NewTask(id, text, true))
		}
	}

	return tasks, nil
}

func (s *SQLiteStorage) UpdateTask(id int, newText string) error {
	_, err := s.DB.Exec("UPDATE tasks SET text = ? WHERE id = ?", newText, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLiteStorage) MarkTask(id int) error {
	_, err := s.DB.Exec("UPDATE tasks SET done = 1 WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLiteStorage) DeleteTask(id int) error {
	_, err := s.DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}