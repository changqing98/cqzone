package service

import (
    "github.com/changqing98/cqzone/user/application/command/protocol"
    userDomain "github.com/changqing98/cqzone/user/domain/user"
    "github.com/changqing98/cqzone/user/infrastructure/constant"
    "github.com/changqing98/cqzone/user/infrastructure/utils/sha"
)

// UserApplicationService 用户应用服务
type UserApplicationService struct {
    UserRepository userDomain.UserRepository
}

// Register 用户注册
func (service *UserApplicationService) Register(command *protocol.CreateUserCommand) *protocol.AccessTokenResponse {
    repo := service.UserRepository
    encryptedPasswd := sha.EncryptWithSalt(command.Password)
    user := &userDomain.User{
        Nickname: command.Email,
        Password: encryptedPasswd,
        Email:    command.Email,
    }
    repo.Save(user)
    return &protocol.AccessTokenResponse{
    }
}

// PasswordLogin 账号密码登录
func (service *UserApplicationService) PasswordLogin(command *protocol.PasswordLoginCommand) *protocol.AccessTokenResponse {
    repo := service.UserRepository
    user := repo.FindByEmail(command.Email)
    if !sha.Verify(user.Password, command.Password) {
        panic(constant.WrongPassword)
    }
    return &protocol.AccessTokenResponse{
    }
}
