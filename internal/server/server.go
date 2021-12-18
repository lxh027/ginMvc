package server

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"io"
	"mvc/internal/global"
	"mvc/tools/formatter"
	"os"
	"time"
)

func Run() {
	serverConfig := global.ServerCfg
	sessionConfig := global.SessionCfg
	// 运行模式
	gin.SetMode(serverConfig.Mode)
	httpServer := gin.Default()

	// 创建session存储引擎
	sessionStore := cookie.NewStore([]byte(sessionConfig.Key))
	sessionStore.Options(sessions.Options{
		MaxAge: sessionConfig.Age,
		Path:   sessionConfig.Path,
	})
	//设置session中间件
	httpServer.Use(sessions.Sessions(sessionConfig.Name, sessionStore))

	gin.DisableConsoleColor()
	// 生成日志
	logPath := fmt.Sprintf("%s-%s.log", global.LogCfg.Path, time.Now().String())
	logFile, _ := os.Create(logPath)
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	// 设置日志格式
	httpServer.Use(gin.LoggerWithFormatter(formatter.GetLogFormat))
	httpServer.Use(gin.Recovery())

	// 注册路由
	ApiRoutes(httpServer)
	BackendRoutes(httpServer)

	serverError := httpServer.Run(serverConfig.Host + ":" + serverConfig.Port)

	if serverError != nil {
		panic("server error !" + serverError.Error())
	}
}
