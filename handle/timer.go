package handle

import (
	"fmt"
	"thewall/model"
	"time"
)

var SensorTypeDic = map[int]string{}
var CropTypeDic = map[int]string{}
var SoilTypeDic = map[int]string{}

func LoadTranslateDic() {
	t := time.Tick(10 * time.Second)
	for {
		<-t
		fmt.Println("每隔 1 秒输出一次")
		//设置传感器类型的映射表
		sensorTypeItems, _ := model.SensorTypeModel.GetAllItems()
		for _, v := range sensorTypeItems {
			SensorTypeDic[v.Id] = v.Name
		}
		//设置土地类型的映射表
		soilTypeItems, _ := model.IotSoilTypeModel.GetAllItems()
		for _, v := range soilTypeItems {
			SoilTypeDic[v.Id] = v.Name
		}
		//设置农作物类型映射表
		cropTypeItems, _ := model.CropTypeModel.GetAllItems()
		for _, v := range cropTypeItems {
			CropTypeDic[v.Id] = v.Name
		}

	}
}
