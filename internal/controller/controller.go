package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"traversal-share/internal/service"
)

func DemoPage() {
	//创建默认路由
	router := gin.Default()
	//处理跨域
	router.Use(Cors())
	//处理请求
	router.GET("/test", service.EchoText)
	//启动监听服务
	router.Run(":8080")
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		//请求头部
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			// 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}
		c.Next()
	}
}
