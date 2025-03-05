// health_check.go
package main

import (
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

var isHealthy int32 = 1

func healthHandler(c *gin.Context) {
	if atomic.LoadInt32(&isHealthy) == 1 {
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Write([]byte("OK"))
	} else {
		c.Writer.WriteHeader(http.StatusServiceUnavailable)
	}
}

func setUnhealthy() {
	atomic.StoreInt32(&isHealthy, 0)
}
