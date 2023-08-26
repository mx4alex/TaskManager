package postgres_storage

import (
	"fmt"
	"TaskManager/internal/entity"
	"TaskManager/internal/config"
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgreSQLStorage struct {
	DB *sql.DB
}

func New(cfg config.PostgresConfig) (*PostgreSQLStorage, error) {
	conn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS tasks (id SERIAL PRIMARY KEY, text TEXT, done INTEGER)")
	if err != nil {
		return nil, err
	}

	return &PostgreSQLStorage{DB: db}, nil
}

func (s *PostgreSQLStorage) AddTask(newText string) error {
	_, err := s.DB.Exec("INSERT INTO tasks (text, done) VALUES ($1, $2)", newText, 0)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgreSQLStorage) GetTasks() ([]entity.Task, error) {
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

func (s *PostgreSQLStorage) UpdateTask(id int, newText string) error {
	_, err := s.DB.Exec("UPDATE tasks SET text = $1 WHERE id = $2", newText, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgreSQLStorage) MarkTask(id int) error {
	_, err := s.DB.Exec("UPDATE tasks SET done = 1 WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgreSQLStorage) DeleteTask(id int) error {
	_, err := s.DB.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
