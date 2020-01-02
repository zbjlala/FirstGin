package main

import (
	"fmt"
	"FirstGin/pkg/setting"
	"FirstGin/routers"
	"net/http"
)

func main() {
	//router := gin.Default()
	//router.GET("/test", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"message": "test",
	//	})
	//})


	//endless.DefaultReadTimeOut = setting.ReadTimeout
	//endless.DefaultWriteTimeOut = setting.WriteTimeout
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	//
	//server := endless.NewServer(endPoint, routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//	log.Println("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil{
	//	log.Printf("Server err: %v", err)
	//}

	router := routers.InitRouter()

	s := &http.Server{
		Addr:	fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:	router,
		ReadTimeout:	setting.ReadTimeout,
		WriteTimeout:	setting.WriteTimeout,
		MaxHeaderBytes:	1 << 20,
	}

	s.ListenAndServe()
}