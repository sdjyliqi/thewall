package utils

import (
	"github.com/golang/glog"
	"strconv"
)

//Convert2Int ... 如果错误，返回值为-1，如果正确
func Convert2Int(src string) int {
	dest, err := strconv.Atoi(src)
	if err != nil {
		glog.Errorf("convert The resource %s to integer failed,err:%+v", src, err)
		return -1
	}
	return dest
}

//Convert2Int64 ... 如果错误，返回值为-1，如果正确
func Convert2Int64(src string) int64 {
	dest, err := strconv.ParseInt(src, 10, 64)
	if err != nil {
		glog.Errorf("convert The resource %s to integer failed,err:%+v", src, err)
		return -1
	}
	return dest
}
