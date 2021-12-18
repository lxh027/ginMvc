package main

import (
	"flag"
	"mvc/internal/global"
	"mvc/internal/server"
)

func main() {
	defer func() {
		_ = global.RedisClient.Client.Close()
		db, _ := global.MysqlClient.DB()
		_ = db.Close()
	}()

	var env string

	flag.StringVar(&env, "env", "dev", "配置环境")
	flag.Parse()

	global.Init(env)
	server.Run()
}
