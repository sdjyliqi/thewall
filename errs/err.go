package errs

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
