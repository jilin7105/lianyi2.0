/*
@Time : 2022/9/22 5:09 PM
@Author : yuyunqing
@File : route
@Software: GoLand
*/
package route

import (
	"github.com/gin-gonic/gin"
	"lianyi/service/UserAbout"
	"lianyi/service/Wifi"
)

func RouteInit(r *gin.Engine) {
	r.GET("/api/checkOpenid", UserAbout.GetUser)          //判断openid是否收录
	r.POST("/api/wifi/set-wifi", Wifi.SetWifi)            //收录wifi
	r.GET("/api/wifi/get-wifi", Wifi.GetWifiById)         //获取wifi根据id ，会用于扫码链接
	r.GET("/api/wifi/get-inv-wifi", Wifi.GetWifiByOpenid) //我收录的wifi
	r.GET("/api/wifi/conn-success", Wifi.WifiConnSucess)  //成功链接记录
	r.GET("/api/wifi/get-wifi-qr", Wifi.GetWifiQr)        //获取二维码
}
