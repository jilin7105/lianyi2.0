/*
@Time : 2022/9/23 10:41 AM
@Author : yuyunqing
@File : WifiConnLog
@Software: GoLand
*/
package model

import "time"

type WifiConnlog struct {
	Id        int32     `gorm:"column:id" json:"id"`
	Openid    string    `gorm:"column:open_id" json:"open_id"`
	WifiId    int64     `gorm:"column:wifi_id" json:"wifi_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (WifiConnlog) TableName() string {
	return "wifi_conn_log"
}
