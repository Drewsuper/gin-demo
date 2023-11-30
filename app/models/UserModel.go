package models

import (
	"errors"
	"fmt"
)

type User struct {
	Id       int    `json:"id" gorm:"column:id;not null autoIncrement;primaryKey"`
	Username string `json:"u_name" gorm:"column:uname"`
	UserPWD  string `json:"pwd" gorm:"column:pwd"`
}

func (User) TableName() string {
	return "b_user"
}

func (u User) ToString() string {
	return fmt.Sprintf("{'id':%d,'uname':%s,'pwd':%s}", u.Id, u.Username, u.UserPWD)
}

func FindOne(name string) User {
	user := User{}
	_, err := DB.Where("uname = ?", name).First(&user).Rows()
	if err != nil {
		return User{}
	}
	return user
}

func FindOneById(id int) User {
	user := User{}
	DB.Where("id = ?", id).First(&user)
	return user
}

func InsertUser(user *User) (int, error) {
	row := TX.Create(user).RowsAffected
	if row < 0 {
		TX.Rollback()
		return -1, errors.New("failed")
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
