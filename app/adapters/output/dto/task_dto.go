package dto

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskDTO struct {
	TaskId          string    `copier:"TaskId" gorm:"primaryKey;column:TaskId"`
	TaskMessage     string    `copier:"TaskMessage" gorm:"column:TaskMessage"`
	CreatedAt       time.Time `copier:"CreatedAt" gorm:"autoCreateTime;column:CreatedAt"`
	IsTaskCompleted bool      `copier:"IsTaskCompleted" default:"false" gorm:"column:IsTaskCompleted"`
	UserId          string    `copier:"UserId" gorm:"column:UserId"`
}

func (TaskDTO) TableName() string {
	return "TASK"
}

func (t *TaskDTO) BeforeCreate(tx *gorm.DB) (err error) {
	t.TaskId = uuid.New().String()
	return
}
