package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//Ping ... 测试接口
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": gin.H{"time": time.Now()}})
}
