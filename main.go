package main

import (
	"FirstGin/models"
	"FirstGin/pkg/logging"
	"FirstGin/pkg/setting"
	"FirstGin/routers"
	"fmt"
	_ "fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_ "net/http"
)

func init(){
	setting.Setup()
	models.Setup()
	logging.Setup()
}

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
	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	s := &http.Server{
		Addr:	endPoint,
		Handler:	routersInit,
		ReadTimeout: readTimeout,
		WriteTimeout:	writeTimeout,
		MaxHeaderBytes:	maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	s.ListenAndServe()
}