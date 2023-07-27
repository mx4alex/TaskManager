package memory_storage_test

import (
	"github.com/stretchr/testify/assert"
	"TaskManager/internal/entity"
	"TaskManager/internal/storage/memory_storage"
	"testing"
)

func TestMemoryStorage(t *testing.T) {
	t.Run("add data", func(t *testing.T) {
		expectedData := []entity.Task {
			entity.NewTask(1, "Task 1", false),
			entity.NewTask(2, "Task 2", false),
		}
	
		ms, _ := memory_storage.New()
		ms.AddTask("Task 1")
		ms.AddTask("Task 2")
	
		resultData, _ := ms.GetTasks()
		assert.Equal(t, expectedData, resultData)
	})

	t.Run("delete data", func(t *testing.T){
		expectedData := []entity.Task {
			entity.NewTask(2, "Task 2", false),
		}
	
		ms, _ := memory_storage.New()
		ms.AddTask("Task 1")
		ms.AddTask("Task 2")
	
		err := ms.DeleteTask(1)
		assert.NoError(t, err)

		resultData, _ := ms.GetTasks()
		assert.Equal(t, expectedData, resultData)
	})

	t.Run("update data", func(t *testing.T){
		expectedData := []entity.Task {
			entity.NewTask(1, "New Task", false),
			entity.NewTask(2, "Task 2", false),
		}
	
		ms, _ := memory_storage.New()
		ms.AddTask("Task 1")
		ms.AddTask("Task 2")
	
		err := ms.UpdateTask(1, "New Task")
		assert.NoError(t, err)

		resultData, _ := ms.GetTasks()
		assert.Equal(t, expectedData, resultData)
	})

	t.Run("test error", func(t *testing.T){
		ms, _ := memory_storage.New()
		ms.AddTask("Task 1")
	
		err := ms.UpdateTask(2, "New Task")
		assert.Error(t, err)
	})

}