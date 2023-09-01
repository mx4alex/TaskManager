package usecase_test

import (
	"testing"
	"TaskManager/internal/entity"
	"github.com/golang/mock/gomock"
	mock "TaskManager/internal/usecase/mock"
	"TaskManager/internal/usecase"
	"errors"
)

func TestTaskInteractor_AddTask(t *testing.T) {
	testCases := []struct {
		name           string
		inputText      string
		expectedError  error
	}{
		{
			name:           "Valid input",
			inputText:      "test task",
			expectedError:  nil,
		},
		{
			name:           "Empty input",
			inputText:      "",
			expectedError:  errors.New("empty input"),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskStorage := mock.NewMockTaskStorage(ctrl)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockTaskStorage.EXPECT().AddTask(tc.inputText).Return(tc.expectedError)

			taskInteractor := usecase.NewTaskInteractor(mockTaskStorage)
			err := taskInteractor.AddTask(tc.inputText)

			if err != tc.expectedError {
				t.Errorf("Expected error: %v, but got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestTaskInteractor_GetTasks(t *testing.T) {
	testCases := []struct {
		name              string
		mockReturnData    []entity.Task
		mockReturnError   error
		expectedTaskCount int
		expectedError     error
	}{
		{
			name:              "One task",
			mockReturnData:    []entity.Task{{ID: 1, Text: "task 1", Done: false}},
			mockReturnError:   nil,
			expectedTaskCount: 1,
			expectedError:     nil,
		},
		{
			name:              "Empty tasks",
			mockReturnData:    []entity.Task{},
			mockReturnError:   nil,
			expectedTaskCount: 0,
			expectedError:     nil,
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskStorage := mock.NewMockTaskStorage(ctrl)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockTaskStorage.EXPECT().GetTasks().Return(tc.mockReturnData, tc.mockReturnError)

			taskInteractor := usecase.NewTaskInteractor(mockTaskStorage)
			tasks, err := taskInteractor.GetTasks()

			if err != tc.expectedError {
				t.Errorf("Expected error: %v, but got: %v", tc.expectedError, err)
			}

			if len(tasks) != tc.expectedTaskCount {
				t.Errorf("Expected %d tasks, but got: %d", tc.expectedTaskCount, len(tasks))
			}
		})
	}
}

func TestTaskInteractor_UpdateTask(t *testing.T) {
	testCases := []struct {
		name           string
		taskID         int
		newText        string
		expectedError  error
	}{
		{
			name:           "Valid input",
			taskID:         1,
			newText:        "updated task",
			expectedError:  nil,
		},
		{
			name:           "Empty text",
			taskID:         2,
			newText:        "",
			expectedError:  errors.New("empty input"),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskStorage := mock.NewMockTaskStorage(ctrl)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockTaskStorage.EXPECT().UpdateTask(tc.taskID, tc.newText).Return(tc.expectedError)

			taskInteractor := usecase.NewTaskInteractor(mockTaskStorage)
			err := taskInteractor.UpdateTask(tc.taskID, tc.newText)

			if err != tc.expectedError {
				t.Errorf("Expected error: %v, but got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestTaskInteractor_MarkTask(t *testing.T) {
	testCases := []struct {
		name           string
		taskID         int
		expectedError  error
	}{
		{
			name:           "Valid input",
			taskID:         1,
			expectedError:  nil,
		},
		{
			name:           "Negative ID",
			taskID:         -1,
			expectedError:  errors.New("invalid task ID"),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskStorage := mock.NewMockTaskStorage(ctrl)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockTaskStorage.EXPECT().MarkTask(tc.taskID).Return(tc.expectedError)

			taskInteractor := usecase.NewTaskInteractor(mockTaskStorage)
			err := taskInteractor.MarkTask(tc.taskID)

			if err != tc.expectedError {
				t.Errorf("Expected error: %v, but got: %v", tc.expectedError, err)
			}
		})
	}
}

func TestTaskInteractor_DeleteTask(t *testing.T) {
	testCases := []struct {
		name           string
		taskID         int
		expectedError  error
	}{
		{
			name:           "Valid input",
			taskID:         1,
			expectedError:  nil,
		},
		{
			name:           "Negative ID",
			taskID:         -1,
			expectedError: errors.New("invalid task ID"),
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskStorage := mock.NewMockTaskStorage(ctrl)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockTaskStorage.EXPECT().DeleteTask(tc.taskID).Return(tc.expectedError)

			taskInteractor := usecase.NewTaskInteractor(mockTaskStorage)
			err := taskInteractor.DeleteTask(tc.taskID)

			if err != tc.expectedError {
				t.Errorf("Expected error: %v, but got: %v", tc.expectedError, err)
			}
		})
	}
}