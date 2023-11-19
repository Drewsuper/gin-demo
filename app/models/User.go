package models

import (
	"fmt"
)

type User struct {
	Id       int    `json:"id" gorm:"column:id;not null AUTO_INCREMENT;primaryKey"`
	Username string `json:"u_name" gorm:"column:u_name"`
	UserPWD  string `json:"pwd" gorm:"column:u_pwd"`
}

func (User) TableName() string {
	return "user"
}

func (u User) ToString() string {
	return fmt.Sprintf("{'id':%d,'uname':%s,'pwd':%s}", u.Id, u.Username, u.UserPWD)
}

func FindOne(data string) User {
	user := User{}
	DB.Where("u_name = ?", data).First(&user)
	return user
}

func FindOneById(id int) User {
	user := User{}
	DB.Where("id = ?", id).First(&user)
	return user
}

func InsertUser(user User) (int, error) {
	var err error
	err = TX.Create(user).Error
	if err != nil {
		TX.Rollback()
		return -1, err
	} else {
		TX.Commit()
		return 1, nil
	}
}

func UpdateData(id int, pwd string) (int, error) {
	rowAffected := TX.Where("id = ?", id).Update("pwd", pwd).RowsAffected
	if rowAffected > 0 {
		TX.Commit()
		return 1, nil
	} else {
		TX.Rollback()
		return -1, nil
	}
}
