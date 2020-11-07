package service

import (
    "github.com/changqing98/cqzone/user/application/command"
    "github.com/changqing98/cqzone/user/application/dto"
)

// UserApplicationService 用户应用服务接口定义
type UserApplicationService interface {
    // 发送手机验证码
    SendSmsVerificationCode(sendSmsVerificationCodeCommand command.SendSmsVerificationCodeCommand) string
    // 用户注册
    Register(createUserCommand command.EmailRegisterCommand) dto.AuthenticationDTO
    // 密码登录
    PasswordLogin(passwordLoginCommand command.PasswordLoginCommand) dto.AuthenticationDTO
}
