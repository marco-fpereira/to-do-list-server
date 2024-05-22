package model

import (
	"time"
	//	"github.com/jinzhu/copier"
)

type TaskDomain struct {
	TaskId            string    `copier:"taskId"`
	TaskMessage       string    `copier:"taskMessage"`
	CreatedAt         time.Time `copier:"createdAt"`
	IsTaskUncompleted bool      `copier:"isTaskCompleted" default:"false"`
}
