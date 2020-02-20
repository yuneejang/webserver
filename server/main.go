package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//Step 1. Creates a gin router
	//- 1. 기본
	// Default With the Logger and Recovery middleware already attached
	router := gin.Default()
	//router := gin.New()
	//- 2.
	// Creates a router without any middleware by default
	//router := gin.New()

	//Setp ?. HTML rendering
	//Using LoadHTMLGlob() or LoadHTMLFiles()
	router.LoadHTMLGlob("templates/*.html")

	// v1 := router.Group("v1")
	// {
	// 	//	v1.POST("/login", loginEndpoint)
	// 	//	v1.POST("/submit", submitEndpoint)
	// 	//	v1.POST("/read", readEndpoint)
	// }
	// // Simple group: v2
	// v2 := router.Group("/v2")
	// {
	// 	//	v2.POST("/login", loginEndpoint)
	// 	//	v2.POST("/submit", submitEndpoint)
	// 	//	v2.POST("/read", readEndpoint)
	// }

	// Using GET, POST, PUT, PATCH, DELETE and OPTIONS
	// router.GET("/welcome", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{ // Go로 HTMl 자주 뿌리진 않겟지만 쓰려면 이렇게.. view에서 읽을 땐 html 참고, 템플릿 엔진은 핸들바랑 비슷 (.빼고)
	// 		"message": "Golang",
	// 	})
	// })
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

	//실행
	router.Run(":8131")
	//세세한 http config설정이필요한 경우 아래와 같이 사용이 가능함
	// s := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        router,
	// 	ReadTimeout:    10 * time.Second,
	// 	WriteTimeout:   10 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	// s.ListenAndServe()
}
