package main

import (
	"github.com/gin-gonic/gin"
	"Cas/config"
	"Cas/controller"
)


func main(){
	//配置文档
	router := gin.Default()
	//全局跨域1
	router.Use(config.Logger())   //使用Logger中间件，router的类型是*gin.Context
	controller.DefaultRoute(router)
	router.Run(":8081")






	//time.Sleep(time.Second*10)
/*	problems := agent.CheckDarklinks()
	for _,p := range problems{
		fmt.Println("the targgted rule is",p)
	}*/
}


