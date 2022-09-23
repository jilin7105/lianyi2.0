/*
@Time : 2022/9/23 9:55 AM
@Author : yuyunqing
@File : Wifi
@Software: GoLand
*/
package model

import "time"

type Wifi struct {
	Id        int32     `gorm:"column:id" json:"id"`
	SSID      string    `gorm:"column:ssid" json:"SSID"`
	BSSID     string    `gorm:"column:bssid" json:"BSSID"`
	Secure    int32     `gorm:"column:secure" json:"Secure"`
	Shopname  string    `gorm:"column:shopname" json:"Shopname"`
	OpenId    string    `gorm:"column:open_id" json:"OpenId"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Password  string    `gorm:"column:password" json:"Password"`
}

func (Wifi) TableName() string {
	return "wifi"
}

//CREATE TABLE `wifi` (
//  `id` int(11) NOT NULL AUTO_INCREMENT,
//  `ssid` varchar(45) DEFAULT NULL COMMENT 'wifi名称',
//  `bssid` varchar(45) DEFAULT NULL COMMENT 'wifi编码',
//  `secure` tinyint(4) DEFAULT NULL COMMENT '是否安全',
//  `shopname` varchar(45) DEFAULT NULL COMMENT '商户名称',
//  `created_at` datetime DEFAULT NULL,
//  `updated_at` datetime DEFAULT NULL,
//  `open_id` varchar(45) DEFAULT NULL COMMENT '邀请者open_id',
//  `password` varchar(100) NOT NULL COMMENT 'wifi密码',
//  PRIMARY KEY (`id`),
//  KEY `idx_bssid` (`bssid`),
//  KEY `idx_openid` (`open_id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8;
