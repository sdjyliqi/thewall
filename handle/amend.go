package handle

import (
	"email-center/model"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"strconv"
)

//GetAmendWords ... 按页获取
func HelloWord(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": gin.H{"name": "liqi ", "amount": 0}})
}

//SearchAmendWords ...
func SearchAmendWords(c *gin.Context) {
	keywords, _ := c.GetQuery("idx")
	if keywords == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": "bad request"})
	}
	items, err := model.AmendModel.SearchItemsByIdx(keywords)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
}

//DelAmendWords ... 获取异常的分行组织机构
func DelAmendWords(c *gin.Context) {
	strID, ok := c.GetQuery("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "id invalid", "data": nil})
		return
	}
	id, err := strconv.Atoi(strID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	err = model.AmendModel.DelItemByID(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": ""})
}

//UpInsertAmendWords ... 获取异常的分行组织机构
func UpInsertAmendWords(c *gin.Context) {
	reqJson := &model.Amend{}
	err := c.ShouldBindJSON(reqJson)
	if err != nil {
		glog.Errorf("The request %+v is invalid,please check.", c.Request)
		c.JSON(http.StatusBadRequest, gin.H{"code": 99999, "msg": "bind json failed."})
		return
	}
	err = model.AmendModel.UpdateItemByID(reqJson)
	if err != nil {
		glog.Errorf("The request %+v is invalid,please check.", c.Request)
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "msg": "update or insert failed." + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": ""})
}
