package main

import (
	"context"
	"fmt"
	"go-figure-bed/pkg/gmysql"
	"go-figure-bed/pkg/logging"
	"go-figure-bed/pkg/setting"
	"go-figure-bed/routers"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//系统初始化
func init() {
	//全局设置
	setting.Setup()
	//日志配置
	logging.Setup()
	//mysql配置
	gmysql.Setup()
}

//系统启动项
func main() {
	//初始化路由
	routersInit := routers.InitRouter()

	//端口
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	//组装服务
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	//启动服务
	logging.AppLogger.Info("start http server listening.", zap.String("endPoint", endPoint))
	go func() {
		if err := server.ListenAndServe(); err != nil {
			logging.AppLogger.Error("Listen fail.", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logging.AppLogger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		//记录日志并退出系统
		logging.AppLogger.Fatal("Server Shutdown:", zap.Error(err))

	}

	logging.AppLogger.Info("Server exiting")
}
