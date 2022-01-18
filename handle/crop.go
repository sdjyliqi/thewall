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

//GetCropTypeAllItems ... 获取crop_type全量数据
func GetCropTypeAllItems(c *gin.Context) {
	type CropTypeShort struct {
		Id   int    `json:"id" "`
		Name string `json:"name" "`
	}
	var showItmes []*CropTypeShort
	//keywords, _ := c.GetQuery("idx")
	//if keywords == "" {
	//	c.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": "bad request"})
	//}
	items, err := model.CropTypeModel.GetAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	for _, v := range items {
		node := &CropTypeShort{
			Id:   v.Id,
			Name: v.Name,
		}
		showItmes = append(showItmes, node)
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": showItmes})
	return
}

//GetCropTypeItemsByPage ... 分页获取crop_type数据
func GetCropTypeItemsByPage(c *gin.Context) {
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
	items, err := model.CropTypeModel.GetItemsByPage(pageId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
	return
}
