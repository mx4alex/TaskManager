package sqlite_storage_test

import (
	"testing"
	"TaskManager/internal/storage/sqlite_storage"
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
)

func TestSQLiteStorage_AddTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sq := sqlite_storage.SQLiteStorage{DB: db}

	mock.ExpectExec("INSERT INTO tasks").
		WithArgs("test task", 0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = sq.AddTask("test task")
	if err != nil {
		t.Errorf("error was not expected while adding task: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSQLiteStorage_GetTasks(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sq := sqlite_storage.SQLiteStorage{DB: db}

	rows := sqlmock.NewRows([]string{"id", "text", "done"}).
		AddRow(1, "task 1", 0).
		AddRow(2, "task 2", 1)

	mock.ExpectQuery("SELECT id, text, done FROM tasks").
		WillReturnRows(rows)

	tasks, err := sq.GetTasks()
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

func TestSQLiteStorage_UpdateTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sq := sqlite_storage.SQLiteStorage{DB: db}

	mock.ExpectExec("UPDATE tasks SET text = (.?) WHERE id = (.?)").
		WithArgs("updated task", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = sq.UpdateTask(1, "updated task")
	if err != nil {
		t.Errorf("error was not expected while updating task: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSQLiteStorage_MarkTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sq := sqlite_storage.SQLiteStorage{DB: db}

	mock.ExpectExec("UPDATE tasks SET done = 1 WHERE id = (.?)").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = sq.MarkTask(1)
	if err != nil {
		t.Errorf("error was not expected while marking task: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestSQLiteStorage_DeleteTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sq := sqlite_storage.SQLiteStorage{DB: db}

	mock.ExpectExec("DELETE FROM tasks WHERE id = (.?)").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = sq.DeleteTask(1)
	if err != nil {
		t.Errorf("error was not expected while deleting task: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
