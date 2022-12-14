/*
@Time : 2022/9/22 5:16 PM
@Author : yuyunqing
@File : User
@Software: GoLand
*/
package UserAbout

import (
	"github.com/gin-gonic/gin"
	"lianyi/db"
	"lianyi/db/model"
	"net/http"
)

//用户登陆及注册
func GetUser(c *gin.Context) {
	openid := c.GetHeader("X-WX-OPENID")
	user := model.WxUsers{}
	if openid == "" {
		c.JSON(http.StatusOK, gin.H{"status": 1001, "msg": "未获取到openid"})
		return
	}
	db_base := db.Get()
	db_base.Where("openid", openid).Last(&user)
	//如果没有就创建一个用户
	if user.Id <= 0 {
		user.Openid = openid
		db_base.Save(&user)
	}
	c.JSON(http.StatusOK, gin.H{"status": 0,
		"data": gin.H{
			"id":   user.Id,
			"type": user.Type,
		},
	})
}
