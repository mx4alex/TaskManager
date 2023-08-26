package postgres_storage_test

import (
	"testing"
	"TaskManager/internal/storage/postgres_storage"
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
)

func TestPostgreSQLStorage_AddTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ps := postgres_storage.PostgreSQLStorage{DB: db}

	mock.ExpectExec("INSERT INTO tasks").
		WithArgs("test task", 0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = ps.AddTask("test task")
	if err != nil {
		t.Errorf("error was not expected while adding task: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPostgreSQLStorage_GetTasks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ps := postgres_storage.PostgreSQLStorage{DB: db}

	rows := sqlmock.NewRows([]string{"id", "text", "done"}).
		AddRow(1, "task 1", 0).
		AddRow(2, "task 2", 1)

	mock.ExpectQuery("SELECT id, text, done FROM tasks").
		WillReturnRows(rows)

	tasks, err := ps.GetTasks()
	if err != nil {
		t.Errorf("error was not expected while getting tasks: %s", err)
	}

	if reflect.DeepEqual(tasks, rows) {
		t.Errorf("tasks are not equal")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPostgreSQLStorage_UpdateTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ps := postgres_storage.PostgreSQLStorage{DB: db}

	mock.ExpectExec("UPDATE tasks SET text = (.+) WHERE id = (.+)").
		WithArgs("updated task", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = ps.UpdateTask(1, "updated task")
	if err != nil {
		t.Errorf("error was not expected while updating task: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPostgreSQLStorage_MarkTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ps := postgres_storage.PostgreSQLStorage{DB: db}

	mock.ExpectExec("UPDATE tasks SET done = 1 WHERE id = (.+)").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = ps.MarkTask(1)
	if err != nil {
		t.Errorf("error was not expected while marking task: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestPostgreSQLStorage_DeleteTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	ps := postgres_storage.PostgreSQLStorage{DB: db}

	mock.ExpectExec("DELETE FROM tasks WHERE id = (.+)").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = ps.DeleteTask(1)
	if err != nil {
		t.Errorf("error was not expected while deleting task: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
