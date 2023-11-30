package models

type Blog struct {
	Id    int    `json:"id" gorm:"column:id; not null AUTO_INCREMENT;primaryKEY"`
	BName string `json:"b_name" gorm:"column:b_name"`
}
