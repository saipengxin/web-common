package web

import (
	"code.aliyun.com/module-go/ilog"
	"github.com/gin-gonic/gin"
	"time"
)

//接收到参数前中间件
func GetMonitor() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		cost := time.Since(t).Seconds() * 1000
		httpCode := c.Writer.Status()
		status := ""
		if httpCode == 200 {
			status = "success"
		} else {
			status = "fail"
		}
		data := map[string]interface{}{
			"class":        "web",
			"host":         c.Request.Host,
			"uri":          c.Request.URL.Path,
			"cost":         cost,
			"requestCount": 1,
			"status":       status,
		}
		//ilog.LogData(ilog.LL_INFO, "WEB-MONITOR", data)
		ilog.LogData(ilog.LL_INFO, "BASE_STATS_DATA", data)
	}
}
