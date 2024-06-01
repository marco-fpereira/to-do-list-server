package output

import (
	"context"
	"to-do-list-server/app/adapters/output/dto"
	"to-do-list-server/app/domain/model"
	"to-do-list-server/app/domain/model/exception"
	"to-do-list-server/app/domain/port/output"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type mysqlDatabaseAdapter struct {
	db *gorm.DB
}

func NewMysqlDatabaseAdapter(
	DB *gorm.DB,
) output.DatabasePort {
	return &mysqlDatabaseAdapter{
		db: DB,
	}
}

func (m *mysqlDatabaseAdapter) GetUser(
	ctx context.Context,
	userId string,
) (*model.UserCredentialsDomain, error) {
	var userCredentialsDTO dto.UserCredentialsDTO
	userCredentialsDTO.UserId = userId

	result := m.db.First(&userCredentialsDTO)
	if result.Error != nil {
		return nil, exception.BuildSqlException(result.Error)
	}

	var user model.UserCredentialsDomain
	copier.Copy(&user, &userCredentialsDTO)

	return &user, nil
}

func (m *mysqlDatabaseAdapter) GetUserByUsername(
	ctx context.Context,
	username string,
) (*model.UserCredentialsDomain, error) {
	var userCredentialsDTO dto.UserCredentialsDTO
	userCredentialsDTO.Username = username

	result := m.db.First(&userCredentialsDTO)
	if result.Error != nil {
		return nil, exception.BuildSqlException(result.Error)
	}

	var user model.UserCredentialsDomain
	copier.Copy(&user, &userCredentialsDTO)

	return &user, nil
}

func (m *mysqlDatabaseAdapter) CreateUser(
	ctx context.Context,
	username string,
	password string,
) error {
	return nil
}

func (m *mysqlDatabaseAdapter) GetAllTasks(
	ctx context.Context,
	userId string,
) (*[]model.TaskDomain, error) {
	var taskDTOSlice []dto.TaskDTO

	result := m.db.Where("UserId = ?", userId).Find(&taskDTOSlice)
	if result.Error != nil {
		return nil, exception.BuildSqlException(result.Error)
	}

	var taskSlice []model.TaskDomain
	copier.Copy(&taskSlice, &taskDTOSlice)
	return &taskSlice, nil
}

func (m *mysqlDatabaseAdapter) CreateTask(
	ctx context.Context,
	userId string,
	taskMessage string,
	isTaskCompleted bool,
) (*model.TaskDomain, error) {
	taskDTO := dto.TaskDTO{
		TaskMessage:     taskMessage,
		IsTaskCompleted: isTaskCompleted,
		UserId:          userId,
	}

	result := m.db.Create(&taskDTO)

	if result.Error != nil {
		return nil, exception.BuildSqlException(result.Error)
	}

	var task model.TaskDomain
	copier.Copy(&task, &taskDTO)
	return &task, nil
}

func (m *mysqlDatabaseAdapter) GetTask(
	ctx context.Context,
	taskId string,
) (*model.TaskDomain, error) {
	taskDTO := dto.TaskDTO{TaskId: taskId}

	result := m.db.First(&taskDTO)
	if result.Error != nil {
		return nil, exception.BuildSqlException(result.Error)
	}

	var task model.TaskDomain
	copier.Copy(&task, &taskDTO)
	return &task, nil
}

func (m *mysqlDatabaseAdapter) UpdateTaskMessage(
	ctx context.Context,
	taskId string,
	taskMessage string,
) error {
	result := m.db.Model(&dto.TaskDTO{}).
		Where("TaskId = ?", taskId).
		UpdateColumn("TaskMessage", taskMessage)

	if result.Error != nil {
		return exception.BuildSqlException(result.Error)
	}

	return nil
}

func (m *mysqlDatabaseAdapter) UpdateTaskCompleteness(
	ctx context.Context,
	taskId string,
	newTaskCompleteness bool,
) error {
	result := m.db.Model(&dto.TaskDTO{}).
		Where("TaskId = ?", taskId).
		UpdateColumn("IsTaskCompleted", newTaskCompleteness)

	if result.Error != nil {
		return exception.BuildSqlException(result.Error)
	}

	return nil
}

func (m *mysqlDatabaseAdapter) DeleteTask(
	ctx context.Context,
	taskId string,
) error {
	result := m.db.Delete(&dto.TaskDTO{TaskId: taskId})
	if result.Error != nil {
		return exception.BuildSqlException(result.Error)
	}
	return nil
}
