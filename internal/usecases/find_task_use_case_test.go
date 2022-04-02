package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/fgmaia/task/internal/domain/entities"
	"github.com/fgmaia/task/internal/usecases"
	"github.com/fgmaia/task/mocks"
	"github.com/fgmaia/task/sample"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindTask(t *testing.T) {
	t.Parallel()

	userTec1 := sample.NewUserEntityRole(entities.ROLE_TECHNICIAN)
	task1 := sample.NewTaskEntityWithUser(*userTec1)

	taskRepositoryMock := &mocks.TaskRepository{}
	taskRepositoryMock.On("FindTask", mock.Anything, task1.ID).Return(task1, nil)

	t.Run("when taskId is invalid should return an error", func(t *testing.T) {
		taskNoId := ""
		taskInvalidId := "invalid-uuid"

		userRepositoryMock := &mocks.UserRepository{}
		userRepositoryMock.On("FindById", mock.Anything, userTec1.ID).Return(&userTec1, nil)

		findTaskUseCase := usecases.NewFindTaskUseCase(userRepositoryMock, taskRepositoryMock)

		output, err := findTaskUseCase.Execute(context.Background(), userTec1.ID, taskNoId)
		assert.Error(t, err)
		assert.Nil(t, output)

		output, err = findTaskUseCase.Execute(context.Background(), userTec1.ID, taskInvalidId)
		assert.Error(t, err)
		assert.Nil(t, output)
	})

	t.Run("when userId is invalid should return an error", func(t *testing.T) {
		userNoId := ""
		userInvalidId := "invalid-uuid"

		userRepositoryMock := &mocks.UserRepository{}
		userRepositoryMock.On("FindById", mock.Anything, userTec1.ID).Return(&userTec1, nil)

		findTaskUseCase := usecases.NewFindTaskUseCase(userRepositoryMock, taskRepositoryMock)

		output, err := findTaskUseCase.Execute(context.Background(), userNoId, task1.ID)
		assert.Error(t, err)
		assert.Nil(t, output)

		output, err = findTaskUseCase.Execute(context.Background(), userInvalidId, task1.ID)
		assert.Error(t, err)
		assert.Nil(t, output)
	})

	t.Run("when not found user should return an error", func(t *testing.T) {
		errUserNotFound := errors.New("erro when try to find user")

		userRepositoryMock := &mocks.UserRepository{}
		userRepositoryMock.On("FindById", mock.Anything, userTec1.ID).Return(nil, errUserNotFound)

		findTaskUseCase := usecases.NewFindTaskUseCase(userRepositoryMock, taskRepositoryMock)

		output, err := findTaskUseCase.Execute(context.Background(), userTec1.ID, task1.ID)
		assert.Error(t, err)
		assert.Nil(t, output)
	})

	t.Run("when finding task error", func(t *testing.T) {
		errTaskNotFound := errors.New("erro when try to find task")

		userRepositoryMock := &mocks.UserRepository{}
		userRepositoryMock.On("FindById", mock.Anything, userTec1.ID).Return(userTec1, nil)

		taskRepositoryMock := &mocks.TaskRepository{}
		taskRepositoryMock.On("FindTask", mock.Anything, task1.ID).Return(nil, errTaskNotFound)

		findTaskUseCase := usecases.NewFindTaskUseCase(userRepositoryMock, taskRepositoryMock)

		output, err := findTaskUseCase.Execute(context.Background(), userTec1.ID, task1.ID)
		assert.Error(t, err)
		assert.Nil(t, output)
	})

	t.Run("when user role is equals to TECHNICIAN and user is not the owner of it", func(t *testing.T) {
		userTec2 := sample.NewUserEntityRole(entities.ROLE_TECHNICIAN)

		userRepositoryMock := &mocks.UserRepository{}
		userRepositoryMock.On("FindById", mock.Anything, userTec2.ID).Return(userTec2, nil)

		findTaskUseCase := usecases.NewFindTaskUseCase(userRepositoryMock, taskRepositoryMock)

		output, err := findTaskUseCase.Execute(context.Background(), userTec2.ID, task1.ID)
		assert.Error(t, err)
		assert.Nil(t, output)
	})

	t.Run("when successfully FindTask", func(t *testing.T) {
		userRepositoryMock := &mocks.UserRepository{}
		userRepositoryMock.On("FindById", mock.Anything, userTec1.ID).Return(userTec1, nil)

		findTaskUseCase := usecases.NewFindTaskUseCase(userRepositoryMock, taskRepositoryMock)

		output, err := findTaskUseCase.Execute(context.Background(), userTec1.ID, task1.ID)
		assert.NoError(t, err)
		assert.NotNil(t, output)
	})

}