package constant

import "github.com/changqing98/cqzone/user/infrastructure/utils/json"

type Error struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

var (
    WrongPassword = &Error{Code: 40001, Message: "密码错误"}
    UserNotFound  = &Error{Code: 40401, Message: "用户不存在"}
    InternalError = &Error{Code: 50001, Message: "内部错误"}
)

func (e *Error) Error() string {
    return json.ToJsonString(e)
}
