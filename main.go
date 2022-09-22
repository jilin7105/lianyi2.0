package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lianyi/db"
	"lianyi/route"
)

func main() {
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}
	r := gin.Default()
	route.RouteInit(r)
	r.Run(":80")
}
