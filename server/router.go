package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuneejang/webserver/rpc"
)

func SetupRouter(router *gin.Engine) *gin.Engine {

	//Setp. HTML rendering
	//Using LoadHTMLGlob() or LoadHTMLFiles()
	router.LoadHTMLGlob("../../templates/*.html") //경로 지정 다시!!!
	router = setRouterDefault(router)

	//wild card :page를 사용하면서 /v1도 그곳에 포함.. 라우터 그룹하려니 아래와같은 문제가 생김..
	//panic: 'v1' in new path '/v1/login' conflicts with existing wildcard ':page' in exist

	//setting router group
	// v1 := router.Group("/v1")
	// setRouterV1(v1)

	// v2 := router.Group("/v2")
	// setRouterV2(v2)

	////////////////////////////////////////////
	//Json Test
	//router = setRouterJSON(router)

	return router
}

func setRouterJSON(router *gin.Engine) *gin.Engine {
	// Using GET, POST, PUT, PATCH, DELETE and OPTIONS
	router.GET("/", rpc.RequestJSON)
	//router.GET("/", rpc.ParseRequest())
	return router
}

func setRouterDefault(router *gin.Engine) *gin.Engine {

	// Using GET, POST, PUT, PATCH, DELETE and OPTIONS
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil) //세번째인자 무엇인지 모르겠음
	})
	router.GET("/:page", func(c *gin.Context) {
		page := c.Param("page")
		c.HTML(http.StatusOK, page, nil) //세번째인자 무엇인지 모르겠음
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
	fmt.Println("CALL ! V2 !!")

	name := c.Param("name")
	c.String(http.StatusOK, "Hello V2 %s", name)
}
