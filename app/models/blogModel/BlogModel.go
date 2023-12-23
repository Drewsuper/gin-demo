package blogModel

import (
	"errors"
	"gin-new/app/models"
	"log"
)

type BlogModel struct {
	Id       int    `json:"id" gorm:"column:id; not null AUTO_INCREMENT;primaryKEY"`
	BName    string `json:"b_name" gorm:"column:b_name"`
	BDate    string `json:"b_date" gorm:"type:date;colum:b_date"`
	BContext string `json:"b_context" gorm:"type:text;column:b_context"`
	BUp      int    `json:"up" gorm:"column:b_up"`
	UpDate   string `json:"update_time" gorm:"column:up_date"`
	UId      int    `json:"u_id" gorm:"column:u_id;ForeignKey:UserModel;AssociationForeignKey:Id"`
}

func (BlogModel) TableName() string {
	return "b_context"
}

func FindOneById(id int) (BlogModel, error) {
	var blog BlogModel
	tx := models.DB.Begin()
	_, err := tx.Where("id = ?", id).First(&blog).Rows()
	tx.Commit()
	if err == nil {
		return blog, nil
	}
	log.Printf("\033[0;31m%v\033[0m", err)
	return BlogModel{}, err
}

func NewBlogContext(blog *BlogModel) (int64, error) {
	tx := models.DB.Begin()
	row := tx.Create(&blog).RowsAffected
	if row > 0 {
		tx.Commit()
		return row, nil
	} else {
		tx.Rollback()
		return -1, errors.New("插入失败")
	}
}

func FindAllBlog(data *[]BlogModel) error {
	tx := models.DB.Begin()
	_, err := tx.Find(&data).Rows()
	return err
}