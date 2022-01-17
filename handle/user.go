package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"math/rand"
	"net/http"
	"strings"
	"thewall/errs"
	"thewall/model"
	"thewall/utils"
	"time"
)

type UserDto struct {
	Email    string `json:"email"`    //邮件
	Nickname string `json:"nickname"` //昵称
	Password string `json:"password"` //密码
	Country  string `json:"country"`  //国家
	Code     string `json:"code"`     //验证码
}

//UCLogin ... 用户登录
func UCLogin(c *gin.Context) {
	user := model.IotUc{}
	bindErr := c.BindJSON(&user)
	if bindErr != nil || (user.Email == "" || user.Password == "") {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	_, err := model.UCModel.Login(user.Email, user.Password)
	if err != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	//通过校验后，需要重新生成token，更新到db中，后续需要写到redis中，为了调试方便，临时token 先不变化
	token := utils.EncodingPassword(user.Email)
	err = model.UCModel.UpdateToken(user.Email, token)
	if err != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": gin.H{"token": token}})
	return
}

//UCRegister ... 用户注册
func UCRegister(c *gin.Context) {
	userDto := UserDto{}
	bindErr := c.BindJSON(&userDto)
	if bindErr != nil || (userDto.Email == "" || userDto.Password == "" || userDto.Nickname == "" || userDto.Code == "" || userDto.Country == "") {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	////获取缓存验证码
	redisCode, redisErr := utils.GetRedisClient().Get(c, userDto.Email).Result()
	if redisErr != nil {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrCodeNotExisted.Code, "msg": errs.ErrCodeNotExisted.MessageEN, "data": nil})
		return
	}
	//验证码校验,为了方便调试增加一个后门吧
	if userDto.Code != "FFF0" && redisCode != userDto.Code {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrCode.Code, "msg": errs.ErrCode.MessageEN, "data": nil})
		_, redisErr := utils.GetRedisClient().Del(c, userDto.Email).Result()
		if redisErr != nil {
			glog.Fatalf("Del redis key:%+v,err:%+v", userDto.Email, redisErr)
		}
		return
	}
	user := model.IotUc{}
	user.Email = userDto.Email
	user.Nickname = userDto.Nickname
	user.Password = utils.EncodingPassword(userDto.Password)
	ok, err := model.UCModel.Register(user)
	if err != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": ok})
	return
}

//UCResetPassword ... 重置密码
func UCResetPassword(c *gin.Context) {
	userDto := UserDto{}
	bindErr := c.BindJSON(&userDto)
	if bindErr != nil || (userDto.Email == "" || userDto.Password == "" || userDto.Code == "") {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	//获取缓存验证码
	redisCode, redisErr := utils.GetRedisClient().Get(c, userDto.Email).Result()
	if redisErr != nil {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrCodeNotExisted.Code, "msg": errs.ErrCodeNotExisted.MessageEN, "data": nil})
		return
	}
	//验证码校验
	if redisCode != userDto.Code {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrCode.Code, "msg": errs.ErrCode.MessageEN, "data": nil})
		_, redisErr := utils.GetRedisClient().Del(c, userDto.Email).Result()
		if redisErr != nil {
			glog.Fatalf("Del redis key:%+v,err:%+v", userDto.Email, redisErr)
		}
		return
	}

	user := model.IotUc{}
	user.Email = userDto.Email
	user.Password = utils.EncodingPassword(userDto.Password)
	ok, err := model.UCModel.ResetPassword(user)
	if err != errs.Succ {
		c.JSON(http.StatusOK, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": ok})
	return
}

//UCGetCode ... 获取验证码
func UCGetCode(c *gin.Context) {
	email, ok := c.GetQuery("email")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
		return
	}
	code := GenValidateCode(6)
	//设置缓存验证码，30分钟有效
	_, redisErr := utils.GetRedisClient().Set(c, email, code, 30*time.Minute).Result()
	if redisErr != nil {
		glog.Fatalf("Set redis key:%+v,err:%+v", email, redisErr)
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrCode.Code, "msg": errs.ErrCode.MessageEN, "data": nil})
		return
	}

	sendErr := utils.SendToMail("sdjyliqi@163.com", email, "[TheWell]Verification Code", []byte("Verification Code:"+code))
	if sendErr != nil {
		glog.Fatalf("Send to email:%+v,err:%+v", email, sendErr)
		c.JSON(http.StatusOK, gin.H{"code": errs.ErrSendEmail.Code, "msg": errs.ErrSendEmail.MessageEN, "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": nil})
	return
}

func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}
