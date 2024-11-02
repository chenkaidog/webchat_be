package po

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AccountID string `gorm:"column:account_id"`
	Email     string `gorm:"column:email"`
	Username  string `gorm:"column:username"`
	Password  string `gorm:"column:password"`
	Salt      string `gorm:"column:salt"`
	Status    string `gorm:"column:status"`
}

func (Account) TableName() string {
	return "account"
}

type Model struct {
	gorm.Model
	ModelId     string `gorm:"column:model_id"`
	Platform    string `gorm:"column:platform"`
	Name        string `gorm:"column:name"`
	DisplayName string `gorm:"column:display_name"`
}

func (Model) TableName() string {
	return "model"
}

type AccountModel struct {
	gorm.Model
	RelationId string `gorm:"column:relation_id"`
	AccountId  string `gorm:"column:account_id"`
	ModelId    string `gorm:"column:model_id"`
}

func (AccountModel) TableName() string {
	return "account_model"
}

type LoginRecord struct {
	gorm.Model
	AccountID string `gorm:"column:account_id"`
	Status    string `gorm:"column:status"`
	IP        string `gorm:"column:ip"`
	Device    string `gorm:"column:device"`
}

func (LoginRecord) TableName() string {
	return "login_record"
}
