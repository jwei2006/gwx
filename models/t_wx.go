package models

import (
	"time"
)

type TWx struct {
	Tag            string    `xorm:"not null pk VARCHAR(36)"`
	AppId          string    `xorm:"not null unique VARCHAR(100)"`
	AppSecret      string    `xorm:"not null VARCHAR(100)"`
	Token          string    `xorm:"VARCHAR(100)"`
	EncodingAeskey string    `xorm:"VARCHAR(100)"`
	PayMchId       string    `xorm:"VARCHAR(36)"`
	PayNotifyUrl   string    `xorm:"VARCHAR(200)"`
	PayKey         string    `xorm:"VARCHAR(100)"`
	Status         int       `xorm:"default 0 INT(11)"`
	Operator       string    `xorm:"VARCHAR(36)"`
	OperateTime    time.Time `xorm:"DATETIME"`
}
