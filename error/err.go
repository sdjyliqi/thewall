package error

import "errors"

//ErrInfo ... 定义全局错误
type ErrInfo struct {
	Code      int
	Err       error
	MessageCN string
	MessageEN string
}

func (t ErrInfo) MsgEN() string {
	return t.MessageEN
}

func (t ErrInfo) MsgCN() string {
	return t.MessageCN
}

var (
	ErrUserName  = ErrInfo{40001, errors.New("invalid username"), "用户名不存在，请重试", "user not existed"}
	ErrConnMysql = ErrInfo{50001, errors.New("connect mysql failed"), "数据库服务器开小差，请稍后重试", "connect mysql failed"}
	ErrSendEmail = ErrInfo{50002, errors.New("send email failed"), "发送邮件异常，请稍后重试", "send the email failed"}
)
