package model

type Hello struct {
	Hello  string `gorm:"column:hello"`
	Author string `gorm:"column:author"`
}

func (Hello) TableName() string {
	return "lutasam"
}
