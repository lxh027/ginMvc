package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func DebugPrintf(format string, values ...interface{}) {
	if gin.IsDebugging() {
		if !strings.HasSuffix(format, "\n") {
			format += "\n"
		}
		_, _ = fmt.Fprintf(gin.DefaultWriter, "[GIN-debug] "+format, values...)
	}
}

func Printf(format string, values ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	_, _ = fmt.Fprintf(gin.DefaultWriter, "[GIN-debug] "+format, values...)
}
