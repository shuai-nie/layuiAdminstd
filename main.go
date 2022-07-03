package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"layuiAdminstd/global"
	"layuiAdminstd/internal/routers"
	"layuiAdminstd/pkg/setting"
	"log"
	"net/http"
	"time"
)

/*var (
	Logger *logger.Logger
)*/

func init(){
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err %v", err)
	}
}

func main() {
	fmt.Println(global.ServerSetting.RunMode)
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":" + global.ServerSetting.HttpPort,
		Handler: router,
		ReadTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriterTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()


}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriterTimeout *= time.Second

	return nil
}
