package main

import (
	"example/gin-learn/middleware"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	f, _ := os.Create("ginLogging.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router.Use(gin.LoggerWithFormatter(middleware.FormatLogsJson))

	auth := gin.BasicAuth(gin.Accounts{
		"user": "pass",
	})

	admin := router.Group("/admin", auth)
	{
		admin.GET("/getUserData", getUserData)
	}

	router.GET("/Ping", pong)
	router.GET("/getUrlData/:name/:id", getUrlData)
	router.GET("/getData1", middleware.Authenticate, middleware.AddHeader, getData1)
	router.GET("/getData2", getData2)
	router.GET("/getData3", getData3)
	router.Run()

	// http.ListenAndServe(":8080", router)

	// server := &http.Server{
	// 	Addr:         ":8080",
	// 	Handler:      router,
	// 	ReadTimeout:  10 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// }

	// server.ListenAndServe()
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "Pong",
	})
}

func getUrlData(c *gin.Context) {
	name := c.Param("name")
	id := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"data": "This Is Using Param",
		"name": name,
		"age":  id,
	})
}

func getUserData(c *gin.Context) {
	name := c.Query("name")
	id := c.Query("id")
	c.JSON(http.StatusOK, gin.H{
		"data": "User Data",
		"name": name,
		"id":   id,
	})
}

func getData1(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": "This is data 1",
	})
}
func getData2(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": "This is data 2",
	})
}
func getData3(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": "This is data 3",
	})
}
