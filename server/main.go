package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//Step 1. Creates a gin router
	//- 1. 기본
	// Default With the Logger and Recovery middleware already attached
	router := gin.Default()
	//- 2.
	// Creates a router without any middleware by default
	//router := gin.New()

	//Step 2. Setup Router
	router = SetupRouter(router)

	//Step 3. Run Server
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
