package dto

type UserCredentialsDTO struct {
	UserId   string `copier:"UserId" gorm:"primaryKey"`
	Username string `copier:"Username"`
	Password string `copier:"Password"`
}

func (UserCredentialsDTO) TableName() string {
	return "ACCOUNT"
}
