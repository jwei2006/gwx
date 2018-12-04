package models

import (
	"time"
)

type TWxUser struct {
	Openid             string    `xorm:"not null pk VARCHAR(50)"`
	AppId              string    `xorm:"not null index(index_wx_user_appid) VARCHAR(100)"`
	Unionid            string    `xorm:"VARCHAR(50)"`
	Nickname           string    `xorm:"VARCHAR(200)"`
	Sex                int       `xorm:"comment('0未知1男2女') TINYINT(4)"`
	City               string    `xorm:"VARCHAR(100)"`
	Province           string    `xorm:"VARCHAR(100)"`
	Country            string    `xorm:"VARCHAR(100)"`
	Language           string    `xorm:"VARCHAR(20)"`
	Headimgurl         string    `xorm:"VARCHAR(255)"`
	Privilege          []byte    `xorm:"BLOB"`
	Status             int       `xorm:"comment('0未关注，2取消关注，11关注') index(index_wx_user_appid) SMALLINT(6)"`
	LastSubscribeTime  time.Time `xorm:"DATETIME"`
	FirstSubscribeTime time.Time `xorm:"DATETIME"`
	PreOpenid          string    `xorm:"VARCHAR(50)"`
	CreateTime         time.Time `xorm:"DATETIME"`
	UpdateTime         time.Time `xorm:"DATETIME"`
}
