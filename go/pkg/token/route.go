package token

import "github.com/gin-gonic/gin"

// 路由注册
var router = gin.Default()

func RegisterRouter() {
	router.POST("/api/v1/generate/token")
	router.Use(JWT()).GET("/api/v1/get/product/list")
}
