package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BackendRoutes(router *gin.Engine) {

	router.StaticFS("/admin", http.Dir("./web"))
}
