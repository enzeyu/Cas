package config

import "github.com/gin-gonic/gin"

//设置全局跨域头
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//指定所有域名访问
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		//允许后续请求携带认证信息
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//允许请求头字段
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Token, Language, From")
		//允许的请求类型
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		//Next方法仅可以在中间件中使用，它在调用的函数中的链中执行挂起的函数
		c.Next()
	}
}