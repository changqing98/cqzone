package protocol

type CreateUserCommand struct {
    // 邮箱
    Email string
    // 验证码
    Code int
    // 密码
    Password string
}

type PasswordLoginCommand struct {
    // 邮箱
    Email string
    // 密码
    Password string
}
