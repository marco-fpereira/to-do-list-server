package dto

import "time"

type TaskDTO struct {
	TaskId            string    `copier:"taskId" gorm:"primaryKey"`
	TaskMessage       string    `copier:"taskMessage"`
	CreatedAt         time.Time `copier:"createdAt" gorm:"autoCreateTime"`
	IsTaskUncompleted bool      `copier:"isTaskCompleted" default:"false"`
	UserId            string
}
