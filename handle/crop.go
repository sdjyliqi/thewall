package handle

import (
	"email-center/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//GetAmendWords ... 按页获取
func HelloWord(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": gin.H{"name": "liqi ", "amount": 0}})
}

//GetCropAllItems ... 获取crop全量数据
func GetCropAllItems(c *gin.Context) {
	//keywords, _ := c.GetQuery("idx")
	//if keywords == "" {
	//	c.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": "bad request"})
	//}
	items, err := models.IotCropEx.GetAllItems()
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
	items, err := models.IotCropEx.GetItemsByPage(pageId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
	return
}
