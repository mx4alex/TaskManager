package memory_storage_test

import (
	"github.com/stretchr/testify/assert"
	"TaskManager/internal/entity"
	"TaskManager/internal/storage/memory_storage"
	"testing"
)

func TestMemoryStorage(t *testing.T) {
	testCases := []struct {
		name         string
		actions      func(ms *memory_storage.MemoryTaskStorage) error
		expectedData []entity.Task
		expectedErr  bool
	}{
		{
			name: "add data",
			actions: func(ms *memory_storage.MemoryTaskStorage) error {
				ms.AddTask("Task 1")
				ms.AddTask("Task 2")
				return nil
			},
			expectedData: []entity.Task {
				entity.NewTask(1, "Task 1", false),
				entity.NewTask(2, "Task 2", false),
			},
			expectedErr: false,
		},
		{
			name: "delete data",
			actions: func(ms *memory_storage.MemoryTaskStorage) error {
				ms.AddTask("Task 1")
				ms.AddTask("Task 2")
				return ms.DeleteTask(1)
			},
			expectedData: []entity.Task {
				entity.NewTask(2, "Task 2", false),
			},
			expectedErr: false,
		},
		{
			name: "mark data",
			actions: func(ms *memory_storage.MemoryTaskStorage) error {
				ms.AddTask("Task 1")
				return ms.MarkTask(1)
			},
			expectedData: []entity.Task {
				entity.NewTask(1, "Task 1", true),
			},
			expectedErr: false,
		},
		{
			name: "update data",
			actions: func(ms *memory_storage.MemoryTaskStorage) error {
				ms.AddTask("Task 1")
				return ms.UpdateTask(1, "New Task")
			},
			expectedData: []entity.Task {
				entity.NewTask(1, "New Task", false),
			},
			expectedErr: false,
		},
		{
			name: "test error",
			actions: func(ms *memory_storage.MemoryTaskStorage) error {
				ms.AddTask("Task 1")
				return ms.UpdateTask(2, "New Task")
			},
			expectedData: []entity.Task {
				entity.NewTask(1, "Task 1", false),
			},
			expectedErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ms, _ := memory_storage.New()
			err := tc.actions(ms)

			if tc.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			resultData, _ := ms.GetTasks()
			assert.Equal(t, tc.expectedData, resultData)
		})
	}
}
