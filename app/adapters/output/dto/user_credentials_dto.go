package dto

type UserCredentialsDTO struct {
	UserId   string `gorm:"primaryKey"`
	Username string
	Password string
}
