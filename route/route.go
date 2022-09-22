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
)

func RouteInit(r *gin.Engine) {
	r.GET("/api/checkOpenid", UserAbout.GetUser)
}
