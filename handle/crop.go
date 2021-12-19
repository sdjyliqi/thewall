package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"thewall/model"
)

//HelloWord ... 测试
func HelloWord(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": gin.H{"name": "liqi ", "amount": 0}})
}

//GetCropAllItems ... 获取crop全量数据
func GetCropAllItems(c *gin.Context) {
	//keywords, _ := c.GetQuery("idx")
	//if keywords == "" {
	//	c.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": "bad request"})
	//}
	items, err := model.IotCropEx.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
	return
}

//GetCropItemsByPage ... 获取crop全量数据
func GetCropItemsByPage(c *gin.Context) {
	strPage, _ := c.GetQuery("page")
	if strPage == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": "bad request"})
		return
	}
	pageId, err := strconv.Atoi(strPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": "bad request"})
		return
	}
	items, err := model.IotCropEx.GetItemsByPage(pageId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
	return
}
