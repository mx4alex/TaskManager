package sqlite_storage

import (
	"database/sql"
	"TaskManager/internal/entity"
	_ "github.com/mattn/go-sqlite3"
)

const tasksDB = "data/tasks.db"

type SQLiteStorage struct {
	db *sql.DB
	id int
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

	return &SQLiteStorage{db: db}, nil
}

func (s *SQLiteStorage) AddTask(newText string) error {
	statement, err := s.db.Prepare("INSERT INTO tasks (id, text, done) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	s.id++

	_, err = statement.Exec(s.id, newText, false)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLiteStorage) GetTasks() ([]entity.Task, error) {
	rows, err := s.db.Query("SELECT id, text, done FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []entity.Task
	for rows.Next() {
		var id int
		var text string
		var done bool

		err := rows.Scan(&id, &text, &done)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, entity.NewTask(id, text, done))
	}

	return tasks, nil
}

func (s *SQLiteStorage) UpdateTask(id int, newText string) error {
	statement, err := s.db.Prepare("UPDATE tasks SET text = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(newText, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLiteStorage) MarkTask(id int) error {
	statement, err := s.db.Prepare("UPDATE tasks SET done = 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLiteStorage) DeleteTask(id int) error {
	statement, err := s.db.Prepare("DELETE FROM tasks WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}