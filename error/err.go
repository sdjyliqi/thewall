package error

import "errors"

//ErrInfo ... 定义全局错误
type ErrInfo struct {
	Code      int
	Err       error
	MessageCN string
	MessageEN string
}

//GetCode ...  获取错误码
func (e ErrInfo) GetCode() int {
	return e.Code
}

// MsgEN... 获取错误的英文提示
func (e ErrInfo) MsgEN() string {
	return e.MessageEN
}

//MsgCN... 获取错误的中文提示
func (e ErrInfo) MsgCN() string {
	return e.MessageCN
}

var (
	ErrBadRequest = ErrInfo{40000, errors.New("request invalid "), "请求参数非法", "invalid field value in request"}
	ErrUCNoUser   = ErrInfo{40001, errors.New("user not existed"), "用户名不存在", "user not existed"}
	ErrUCPassword = ErrInfo{40002, errors.New("password invalid"), "密码错误", "password invalid"}
	ErrSendEmail  = ErrInfo{40003, errors.New("send email failed"), "发送邮件异常", "send the email failed"}
)
