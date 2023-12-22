package types

import (
	"gin-new/app/models/blogModel"
	"time"
)

const timeLayout = "2006-01-02 15:04:05"

type (
	BlogFindReq struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	BlogNewReq struct {
		BName    string `json:"b_name"`
		BDate    string `json:"b_date"`
		BContext string `json:"b_context"`
		BUpDate  string `json:"up_date"`
		BUp      int    `json:"b_up"`
		UserId   int    `json:"u_id"`
	}
)

func (b BlogNewReq) ToBlogModel() blogModel.BlogModel {
	if len(b.BUpDate) == 0 {
		b.BUpDate = time.Now().Format(timeLayout)
	}
	if len(b.BDate) == 0 {
		b.BDate = time.Now().Format(timeLayout)
	}
	return blogModel.BlogModel{
		BName:    b.BName,
		BDate:    b.BDate,
		BContext: b.BContext,
		BUp:      b.BUp,
		UpDate:   b.BUpDate,
		UId:      b.UserId,
	}
}
