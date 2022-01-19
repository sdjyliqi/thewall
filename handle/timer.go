package handle

import (
	"github.com/golang/glog"
	"thewall/model"
	"time"
)

var SensorTypeDic = map[int]string{}
var CropTypeDic = map[int]string{}
var SoilTypeDic = map[int]string{}

func ConvertCropTypeName(id int) string {
	if id <= 0 {
		return "unknown"
	}
	v, ok := CropTypeDic[id]
	if !ok {
		glog.Errorf("Do not find the crop type %+v,please check it!", id)
	}
	return v
}
func LoadTranslateDic() {
	t := time.Tick(30 * time.Second)
	for {
		<-t
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
