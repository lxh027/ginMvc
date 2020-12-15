package main

import (
	"Authority/server"
	"github.com/gin-gonic/gin"
)
import "Authority/db_server"

var httpServer *gin.Engine

func main()  {
	defer db_server.MySqlDb.Close()

	server.Run(httpServer)

}