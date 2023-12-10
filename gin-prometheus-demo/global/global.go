package global

import "github.com/gin-gonic/gin"

var (
	router *gin.Engine
	dsn    = ""
)

func Dsn() string {
	return dsn
}

func GetRouter() *gin.Engine {
	router = gin.Default()
	return router
}
