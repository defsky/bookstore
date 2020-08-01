package main

import (
	"log"
	"net/http"

	"github.com/defsky/bookstore/user-api/controller"
	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	swgFiles "github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/defsky/bookstore/user-api/docs"
)

// @title Bookstore API
// @version 1.0
// @description This is RESTful API for bookstore project
// @termsOfService http://www.sun-hoo.cn/terms/
// @contact.name API Support
// @contact.url http://www.sun-hoo.cn/surpport
// @contact.email defsky@qq.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 192.168.100.210:8080
// @BasePath /api/v1
func main() {
	accounts := gin.Accounts{
		"def": "123",
	}
	basicAuth := gin.BasicAuth(accounts)

	s := gin.Default()
	authed := s.Group("/admin", basicAuth)

	c := controller.NewController()
	v1 := s.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET(":id", c.GetUser)
			users.POST("", c.AddUser)
		}
	}

	authed.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title": "user-api",
		})
	})

	s.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello,gin world!")
	})

	s.GET("/alive", statusHandler)

	//url := ginSwagger.URL("http://192.168.100.210:8080/swg/doc.json") // The url pointing to API definition
	s.GET("/swagger/*any", ginSwagger.WrapHandler(swgFiles.Handler)) //, url))

	if err := s.Run(); err != nil {
		log.Fatalln(err)
	}
}

func statusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"alive": 1,
	})
}
