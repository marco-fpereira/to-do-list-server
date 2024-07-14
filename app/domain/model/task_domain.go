package model

import (
	"time"
)

type TaskDomain struct {
	TaskId          string    `copier:"TaskId"`
	TaskMessage     string    `copier:"TaskMessage"`
	CreatedAt       time.Time `copier:"CreatedAt"`
	IsTaskCompleted bool      `copier:"IsTaskCompleted" default:"false"`
	UserId          string    `copier:"UserId"`
}
