package dto

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserCredentialsDTO struct {
	UserId   string `copier:"UserId" gorm:"primaryKey;column:UserId"`
	Username string `copier:"Username"`
	Password string `copier:"Password"`
}

func (UserCredentialsDTO) TableName() string {
	return "ACCOUNT"
}

func (u *UserCredentialsDTO) BeforeCreate(tx *gorm.DB) (err error) {
	u.UserId = uuid.New().String()
	return
}
