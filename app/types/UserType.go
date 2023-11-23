package types

type (
	UserLoginReq struct {
		Uname string `json:"username"`
		PWD   string `json:"pwd"`
	}
	UserNewReq struct {
		Uname string `json:"uname"`
		Pwd   string `json:"pwd"`
	}
)
