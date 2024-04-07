package api

import (
	"github.com/gin-gonic/gin"
	"traversal-share/internal/service"
)

func DemoPage() {
	//创建默认路由
	router := gin.Default()
	// 收到请求由对应函数处理
	router.GET("/test", service.EchoText)
	//启动监听服务
	router.Run(":8080")
}
