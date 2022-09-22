/*
@Time : 2022/9/22 5:18 PM
@Author : yuyunqing
@File : WxUsers
@Software: GoLand
*/
package model

import "time"

type WxUsers struct {
	Id        int32     `gorm:"column:id" json:"id"`
	Openid    string    `gorm:"column:openid" json:"openid"`
	Type      int32     `gorm:"column:type" json:"type"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (WxUsers) TableName() string {
	return "wx_user"
}
