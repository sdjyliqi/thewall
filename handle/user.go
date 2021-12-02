package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//UCLogin ... 用户登录
func UCLogin(c *gin.Context) {
	strEmail, _ := c.GetQuery("email")
	strPassport, _ := c.GetQuery("passport")
	fmt.Println(strEmail, strPassport)
	return
}

//UCRegister ... 用户注册
func UCRegister(c *gin.Context) {
	strEmail, _ := c.GetPostForm("email")
	strPassport, _ := c.GetQuery("passport")
	fmt.Println(strEmail, strPassport)
	return
}
