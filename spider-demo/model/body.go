package model

type Body struct {
	Id       string `gorm:"varchar(64); primary_key ;comment: '房子id'"`
	TextBody string `json:"text_body"`
}
