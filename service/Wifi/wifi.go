/*
@Time : 2022/9/23 10:00 AM
@Author : yuyunqing
@File : wifi
@Software: GoLand
*/
package Wifi

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"lianyi/db"
	"lianyi/db/model"
	"log"
	"net/http"
	"strconv"
)

//用户登陆及注册
func SetWifi(c *gin.Context) {
	openid := c.GetHeader("X-WX-OPENID")
	wifi := model.Wifi{}
	c.BindJSON(&wifi)
	if wifi.SSID == "" || wifi.Password == "" {
		c.JSON(http.StatusOK, gin.H{"status": 1002, "msg": "未检测到密码或者wifi信息，请验证后提交"})
		return
	}
	db_base := db.Get()
	if wifi.BSSID != "" {

		var count int64
		db_base.Where(&model.Wifi{BSSID: wifi.BSSID}).Count(&count)
		if count > 0 {
			c.JSON(http.StatusOK, gin.H{"status": 1003, "msg": "此wifi已被收录请检查"})
			return
		}
	}

	wifi.OpenId = openid
	db_base.Save(&wifi)
	c.JSON(http.StatusOK, gin.H{"status": 0,
		"data": gin.H{
			"id": wifi.Id,
		},
	})
}

//获取我邀请的所有wifi
func GetWifiByOpenid(c *gin.Context) {
	openid := c.GetHeader("X-WX-OPENID")
	var wifis []model.Wifi

	db_base := db.Get()
	db_base.Where(&model.Wifi{OpenId: openid}).Find(&wifis)
	if len(wifis) > 0 {
		c.JSON(http.StatusOK, gin.H{"status": 0,
			"data": gin.H{
				"wifis": wifis,
			},
		})
	}

}

//获取wifi密码
func GetWifiById(c *gin.Context) {
	//openid := c.GetHeader("X-WX-OPENID")
	var wifi model.Wifi
	id := c.Query("id")

	if id == "" {
		c.JSON(http.StatusOK, gin.H{"status": 1003, "msg": "系统异常"})
		return
	}
	db_base := db.Get()
	db_base.First(&wifi, id)
	if wifi.Id <= 0 {
		c.JSON(http.StatusOK, gin.H{"status": 1003, "msg": "系统异常"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": 0,
		"data": gin.H{
			"wifi": wifi,
		},
	})

}

//链接成功日志
func WifiConnSucess(c *gin.Context) {
	openid := c.GetHeader("X-WX-OPENID")
	var wifiConnlog model.WifiConnlog
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1003, "msg": ""})
		return
	}
	if atoi <= 0 {
		c.JSON(http.StatusOK, gin.H{"status": 1004, "msg": ""})
		return
	}
	wifiConnlog.WifiId = int64(atoi)
	wifiConnlog.Openid = openid
	db_base := db.Get()
	db_base.Save(&wifiConnlog)

	c.JSON(http.StatusOK, gin.H{"status": 0})

}

//链接成功日志
func GetWifiQr(c *gin.Context) {
	//openid := c.GetHeader("X-WX-OPENID")

	data := make(map[string]interface{})
	id := c.Query("id")
	atoi, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": 1003, "msg": ""})
		return
	}
	if atoi <= 0 {
		c.JSON(http.StatusOK, gin.H{"status": 1004, "msg": ""})
		return
	}

	data["page"] = "pages/connwifi/connwifi"
	data["scene"] = id
	data["check_path"] = false
	data["env_version"] = "release"
	marshal, err := json.Marshal(data)

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://api.weixin.qq.com/wxa/getwxacodeunlimit", bytes.NewBuffer(marshal))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"status": 1004, "msg": "异常"})
		return
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body) //读取数据
	log.Println(body)
	obj := make(map[string]interface{})
	if err := json.Unmarshal(body, &obj); err == nil {
		c.JSON(http.StatusOK, gin.H{"status": 1004, "data": obj})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 0, "data": body})

}
