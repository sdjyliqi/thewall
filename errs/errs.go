package errs

import "errors"

var (
	Succ = ErrInfo{0, nil, "成功", "succ"}

	ErrBadRequest     = ErrInfo{40000, errors.New("request invalid "), "请求参数非法", "invalid field value in request"}
	ErrUCNoUser       = ErrInfo{40001, errors.New("user not existed"), "用户名不存在", "user not existed"}
	ErrUCPassword     = ErrInfo{40002, errors.New("password invalid"), "密码错误", "password invalid"}
	ErrSendEmail      = ErrInfo{40003, errors.New("send email failed"), "发送邮件异常", "send the email failed"}
	ErrUCUserExisted  = ErrInfo{40004, errors.New("user existed"), "用户名已存在", "user existed"}
	ErrCodeNotExisted = ErrInfo{40005, errors.New("code not existed"), "验证码不存在", "code not existed"}
	ErrCode           = ErrInfo{40006, errors.New("code invalid"), "验证码错误", "code invalid"}
	ErrAdd            = ErrInfo{40007, errors.New("add failed"), "添加失败", "add failed"}

	ErrDBGet    = ErrInfo{41000, errors.New("mysql select abnormal"), "数据库查询异常", "database select abnormal"}
	ErrDBInsert = ErrInfo{41001, errors.New("mysql insert abnormal"), "数据库插入异常", "database insert abnormal"}
	ErrDBUpdate = ErrInfo{41002, errors.New("mysql update abnormal"), "数据库更新异常", "database update abnormal"}
	ErrDBDel    = ErrInfo{41003, errors.New("mysql delete abnormal"), "数据库删除异常", "database delete abnormal"}

	//定义生长状态的相关错误
	ErrPeriodNoPlanting = ErrInfo{42000, errors.New("field-not-idle"), "该农田正处于耕种中", "field is using，please ended firstly"}
	ErrPeriodNoHarvest  = ErrInfo{42001, errors.New("field-not-harvest"), "非收割状态", "field-not-harvest"}
	ErrPeriodNoWeigh    = ErrInfo{42002, errors.New("field-not-weigh"), "非称重状态", "dfield-not-weigh"}
)
