package service

import (
    "github.com/changqing98/cqzone/user/application/command/protocol"
    "github.com/changqing98/cqzone/user/domain/user"
)

// UserApplicationService 用户应用服务
type UserApplicationService struct {
    userRepository user.UserRepository
}

// Register 用户注册
func (userApplicationService UserApplicationService) Register(createUserCommand *protocol.CreateUserCommand) *protocol.AuthenticationDTO {
    return &protocol.AuthenticationDTO{
    }
}

// PasswordLogin 账号密码登录
func (userApplicationService UserApplicationService) PasswordLogin(command *protocol.PasswordLoginCommand) *protocol.AuthenticationDTO {
    return &protocol.AuthenticationDTO{
    }
}
