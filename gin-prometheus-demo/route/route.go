package route

import (
	"github.com/zyb-992/gin-prometheus-demo/api/v1/user"
	"github.com/zyb-992/gin-prometheus-demo/global"
	"github.com/zyb-992/gin-prometheus-demo/middleware"
)

func Register() {
	r := global.GetRouter()
	v1 := r.Group("/v1")
	v1.Use(middleware.Jwt())
	{
		v1.POST("/user/register", user.Register)
	}

}
