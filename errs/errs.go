package errs

import "errors"

var (
	Succ = ErrInfo{0, nil, "成功", "succ"}

	ErrBadRequest = ErrInfo{40000, errors.New("request invalid "), "请求参数非法", "invalid field value in request"}
	ErrUCNoUser   = ErrInfo{40001, errors.New("user not existed"), "用户名不存在", "user not existed"}
	ErrUCPassword = ErrInfo{40002, errors.New("password invalid"), "密码错误", "password invalid"}
	ErrSendEmail  = ErrInfo{40003, errors.New("send email failed"), "发送邮件异常", "send the email failed"}

	ErrDBGet    = ErrInfo{41000, errors.New("mysql abnormal"), "数据库查询异常", "database abnormal"}
	ErrDBUpdate = ErrInfo{41002, errors.New("mysql abnormal"), "数据库更新异常", "database abnormal"}
)
