package dto

import "time"

type TaskDTO struct {
	TaskId          string    `copier:"taskId" gorm:"primaryKey" gorm:"column:TaskId"`
	TaskMessage     string    `copier:"taskMessage" gorm:"column:TaskMessage"`
	CreatedAt       time.Time `copier:"createdAt" gorm:"autoCreateTime" gorm:"column:CreatedAt"`
	IsTaskCompleted bool      `copier:"isTaskCompleted" default:"false" gorm:"column:IsTaskCompleted"`
	UserId          string    `copier:"UserId"`
}

func (TaskDTO) TableName() string {
	return "TASK"
}
