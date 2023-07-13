package models

type Category struct {
	Id    int64  `gorm:"primaryKey:autoIncrement"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Icon  string `json:"icon"`
	Color string `json:"color"`
}

func (category *Category) TableName() string {
	return "category"
}
