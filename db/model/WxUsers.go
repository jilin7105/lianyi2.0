/*
@Time : 2022/9/22 5:18 PM
@Author : yuyunqing
@File : WxUsers
@Software: GoLand
*/
package model

import "time"

// CounterModel 计数器模型
type WxUsers struct {
	Id        int32     `gorm:"column:id" json:"id"`
	Openid    string    `gorm:"column:openid" json:"openid"`
	Type      int32     `gorm:"column:type" json:"type"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
}

func (WxUsers) TableName() string {
	return "wx_user"
}
