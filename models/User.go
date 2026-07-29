package models

type User struct {
	UserId   uint   `gorm:"column:UserId;primaryKey;autoIncrement" json:"UserId"`
	Username string `gorm:"column:Username;uniqueIndex" json:"Username"`
	Password string `gorm:"column:Password" json:"Password"`
}

func (User) TableName() string {
	return "user"
}