package models

type UserAuth struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"-"`
	Username string `gorm:"type:varchar(255)" json:"username"`
	Password string `json:"password"`
}
