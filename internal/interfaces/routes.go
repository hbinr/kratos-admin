package interfaces

import (
	"github.com/gin-gonic/gin"
)

func RegisterHTTPServer(
	us *UserUseCase) *gin.Engine {
	router := gin.New()

	rootGrp := router.Group("/api")
	{
		userGrp := rootGrp.Group("/user")
		userGrp.GET("/hi", us.SayHi)
		userGrp.POST("/signup", us.Signup)
		userGrp.PUT("/", us.Update)
		userGrp.GET("/", us.Get)
		userGrp.GET("/list", us.List)
		userGrp.DELETE("/", us.Delete)
	}

	return router
}
