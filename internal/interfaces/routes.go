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
		userGrp.GET("/list", us.List)
		userGrp.GET("/", us.Get)
		userGrp.DELETE("/", us.Delte)
	}

	return router
}
