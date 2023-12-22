package userModel

import (
	"errors"
	"fmt"
	"gin-new/app/models"
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

func FindOne(name string) (User, error) {
	user := User{}
	tx := models.DB.Begin()
	_, err := tx.Where("uname = ?", name).First(&user).Rows()
	tx.Commit()
	if err == nil {
		return user, nil
	}
	return User{}, err
}

func FindOneById(id int) User {
	user := User{}
	models.DB.Where("id = ?", id).First(&user)
	return user
}

func InsertUser(user *User) (int, error) {
	tx := models.DB.Begin()
	row := tx.Create(user).RowsAffected
	if row < 0 {
		tx.Rollback()
		return -1, errors.New("failed")
	} else {
		tx.Commit()
		return 1, nil
	}
}

func UpdateData(id int, pwd string) (int, error) {
	tx := models.DB.Begin()
	rowAffected := tx.Where("id = ?", id).Update("pwd", pwd).RowsAffected
	if rowAffected > 0 {
		tx.Commit()
		return 1, nil
	} else {
		tx.Rollback()
		return -1, nil
	}
}
