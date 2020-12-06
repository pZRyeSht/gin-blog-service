package main

import (
	"gin_blog_service/global"
	"gin_blog_service/internal/model"
	"gin_blog_service/internal/routers"
	"gin_blog_service/pkg/logger"
	setting2 "gin_blog_service/pkg/setting"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

// @title 博客后端
// @version 1.0
// @description Go 语言编程之旅：一起用Go做项目
func main() {
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200,gin.H{"massage":"pong"})
	//})
	//_ = r.Run()
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}

func init() {
	if err := setupSetting(); err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	if err := setupDBEngine(); err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func setupSetting() error {
	setting, err := setting2.NewSetting()
	if err != nil {
		return err
	}
	if err = setting.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}
	if err = setting.ReadSection("App", &global.AppSetting); err != nil {
		return err
	}
	if err = setting.ReadSection("Database", &global.DatabaseSetting); err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}
