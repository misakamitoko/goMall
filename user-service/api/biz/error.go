package biz

const OK = 200

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func NewError(code int, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

func (e *Error) Error() string {
	return e.Msg
}

func Fail(code int, msg string) *Result {
	return &Result{
		Code: code,
		Msg:  msg,
	}
}

func Success(data any) *Result {
	return &Result{
		Code: OK,
		Msg:  "success",
		Data: data,
	}
}

var (
	DbError              = NewError(10000, "数据库错误")
	AlreadyExistError    = NewError(100001, "已存在")
	UserNotExistError    = NewError(100002, "用户不存在")
	TokenError           = NewError(100003, "鉴权不通过")
	RedisError           = NewError(100004, "redis错误")
	NoRegisterError      = NewError(100005, "未注册")
	PasswordNotMathError = NewError(100006, "密码不匹配")
)
