package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) *gin.Engine {

	//Setp. HTML rendering
	//Using LoadHTMLGlob() or LoadHTMLFiles()
	router.LoadHTMLGlob("templates/*.html")
	router = setRouterDefault(router)

	v1 := router.Group("v1")
	setRouterV1(v1)

	v2 := router.Group("/v2")
	setRouterV2(v2)
	return router
}

func setRouterDefault(router *gin.Engine) *gin.Engine {

	// Using GET, POST, PUT, PATCH, DELETE and OPTIONS
	fmt.Println("setRouterDefault")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil) //세번째인자 무엇인지 모르겠음
	})
	router.GET("/charts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "charts.html", nil) //세번째인자 무엇인지 모르겠음
	})
	router.GET("/tables", func(c *gin.Context) {
		c.HTML(http.StatusOK, "tables.html", nil) //세번째인자 무엇인지 모르겠음
	})
	router.GET("/utilities-color", func(c *gin.Context) {
		c.HTML(http.StatusOK, "utilities-color.html", nil) //세번째인자 무엇인지 모르겠음
	})
	router.GET("/utilities-other", func(c *gin.Context) {
		c.HTML(http.StatusOK, "utilities-other.html", nil) //세번째인자 무엇인지 모르겠음
	})
	router.GET("/blank", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blank.html", nil) //세번째인자 무엇인지 모르겠음
	})
	router.GET("/buttons", func(c *gin.Context) {
		c.HTML(http.StatusOK, "buttons.html", nil) //세번째인자 무엇인지 모르겠음
	})
	router.GET("/cards", func(c *gin.Context) {
		c.HTML(http.StatusOK, "cards.html", nil) //세번째인자 무엇인지 모르겠음
	})
	router.GET("/forgot-password", func(c *gin.Context) {
		c.HTML(http.StatusOK, "forgot-password.html", nil) //세번째인자 무엇인지 모르겠음
	})
	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil) //세번째인자 무엇인지 모르겠음
	})
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil) //세번째인자 무엇인지 모르겠음
	})
	router.GET("/utilities-animation", func(c *gin.Context) {
		c.HTML(http.StatusOK, "utilities-animation", nil) //세번째인자 무엇인지 모르겠음
	})
	router.GET("/404", func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.html", nil) //세번째인자 무엇인지 모르겠음
	})
	router.GET("/utilities-border", func(c *gin.Context) {
		c.HTML(http.StatusOK, "utilities-border.html", nil) //세번째인자 무엇인지 모르겠음
	})
	return router
}

func setRouterV1(v1 *gin.RouterGroup) {
	v1.GET("/login", callV1)
	v1.GET("/submit", callV1)
	v1.GET("/read", callV1)
}

func setRouterV2(v2 *gin.RouterGroup) {
	{
		v2.GET("/login", callV2)
		v2.GET("/submit", callV2)
		v2.GET("/read", callV2)
	}

}

func callV1(c *gin.Context) {
	fmt.Println("CALL ! V1 !!")
	c.String(http.StatusOK, "Hello V1 ")

}

func callV2(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello V2 %s", name)
}
