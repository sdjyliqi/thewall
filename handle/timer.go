package handle

import (
	"fmt"
	"github.com/golang/glog"
	"thewall/model"
	"time"
)

var CropTypeDic = map[int]string{}
var ProbeTypeDic = map[int]string{}
var SoilTypeDic = map[int]string{}
var ReferenceDic = map[string]*model.IotReference{}

//GetCropTypeByID
func GetCropTypeByID(id int) string {
	if id <= 0 {
		return "unknown"
	}
	v, ok := CropTypeDic[id]
	if !ok {
		glog.Errorf("Do not find the crop type %+v,please check it!", id)
	}
	return v
}

//GetProbeTypeByID...
func GetProbeTypeByID(id int) string {
	if id <= 0 {
		return "unknown"
	}
	v, ok := ProbeTypeDic[id]
	if !ok {
		glog.Errorf("Do not find the Sensor type %+v,please check it!", id)
	}
	return v
}

//GetSoilTypeByID...
func GetSoilTypeByID(id int) string {
	if id <= 0 {
		return "unknown"
	}
	v, ok := SoilTypeDic[id]
	if !ok {
		glog.Errorf("Do not find the soil type %+v,please check it!", id)
	}
	return v
}

func createIdxForReference(soilTypeID, cropTypeID int) string {
	return fmt.Sprintf("%d_%d", soilTypeID, cropTypeID)
}

//GetReference... 获取
func GetReference(soilTypeID, cropTypeID int) *model.IotReference {
	v, ok := ReferenceDic[createIdxForReference(soilTypeID, cropTypeID)]
	if !ok {
		glog.Errorf("Do not find the item for soil_type_id %d with crop_type_id %d ,please check it!", soilTypeID, cropTypeID)
		return nil
	}
	return v
}

func GetReferenceNotice(soilTypeID, cropTypeID int) string {
	item := GetReference(soilTypeID, cropTypeID)
	if item == nil {
		return "unknown"
	}
	return fmt.Sprintf("%.2f %%----%.2f %%", item.HumidityMin*100, item.HumidityMax*100)
}
func LoadTranslateDic() {
	t := time.Tick(3 * time.Second)
	for {
		<-t
		//设置传感器类型的映射表
		probeTypeItems, _ := model.ProbeTypeModel.GetAllItems()
		for _, v := range probeTypeItems {
			ProbeTypeDic[v.Id] = v.Name
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
		//根据土地类型和种植农作物的类型，查询湿度范围
		referenceItems, _ := model.ReferenceModel.GetAllItems()
		for _, v := range referenceItems {
			ReferenceDic[createIdxForReference(v.SoilTypeId, v.CropTypeId)] = v
		}
		t = time.Tick(300 * time.Second)
	}
}
