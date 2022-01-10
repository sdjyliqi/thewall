package handle

//
////GetProjectAllItems ... 获取Project全量数据
//func GetProjectAllItems(c *gin.Context) {
//	items, err := model.ProjectModel.GetAllItems()
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
//	return
//}
//
////GetProjectItemsByPage ... 分页获取Project全量数据
//func GetProjectItemsByPage(c *gin.Context) {
//	strPage, _ := c.GetQuery("page")
//	if strPage == "" {
//		c.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": "bad request"})
//		return
//	}
//	pageId, err := strconv.Atoi(strPage)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"code": 0, "msg": "bad request"})
//		return
//	}
//	items, err := model.ProjectModel.GetItemsByPage(pageId)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "msg": err.Error(), "data": nil})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
//	return
//}
//
////AddProject ... 添加一条Project数据
//func AddProject(c *gin.Context) {
//	item := model.IotProject{}
//	bindErr := c.BindJSON(&item)
//	if bindErr != nil {
//		c.JSON(http.StatusOK, gin.H{"code": errs.ErrBadRequest.Code, "msg": errs.ErrBadRequest.MessageEN, "data": nil})
//		return
//	}
//	ok, err := model.ProjectModel.AddItem(&item)
//	if err != errs.Succ {
//		c.JSON(http.StatusOK, gin.H{"code": err.Code, "msg": err.MessageEN, "data": nil})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": ok})
//	return
//}
