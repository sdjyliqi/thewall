package handle

import (
	"email-center/errs"
	"email-center/model"
	"email-center/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//UCLogin ... 用户登录
func UCLogin(c *gin.Context) {
	strEmail, _ := c.GetQuery("email")
	strPassword, _ := c.GetQuery("pp")
	if strEmail == "" || strPassword == "" {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	_, err := models.UCModel.Login(strEmail, strPassword)
	if err != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	//通过校验后，需要重新生成token，更新到db中，后续需要写到redis中，为了调试方便，临时token 先不变化
	token := utils.EncodingPassword(strEmail)
	err = models.UCModel.UpdateToken(strEmail, token)
	if err != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": gin.H{"token": token}})
	return
}

//UCRegister ... 用户注册
func UCRegister(c *gin.Context) {
	strEmail, _ := c.GetPostForm("email")
	strPassport, _ := c.GetQuery("passport")
	fmt.Println(strEmail, strPassport)
	return
}
