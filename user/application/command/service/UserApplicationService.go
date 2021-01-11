package service

import (
    "github.com/changqing98/cqzone/user/application/command"
    "github.com/changqing98/cqzone/user/application/dto"
    "github.com/changqing98/cqzone/user/domain/user/model"
    "github.com/changqing98/cqzone/user/domain/user/repository"
    infraSvc "github.com/changqing98/cqzone/user/infrastructure/service"
    "github.com/changqing98/cqzone/user/infrastructure/utils/sha"
    "strconv"
)

// UserApplicationService 用户应用服务
type UserApplicationService struct {
    userRepository repository.UserRepository
    smsService     infraSvc.SmsService
}

func NewUserApplicationService(userRepository repository.UserRepository, smsService infraSvc.SmsService) UserApplicationService {
    return UserApplicationService{userRepository, smsService}
}

// Register 用户注册
func (userApplicationService UserApplicationService) Register(createUserCommand *command.CreateUserCommand) *dto.AuthenticationDTO {
    userRepo := userApplicationService.userRepository
    user := model.User{
        Username: createUserCommand.Username,
        Password: createUserCommand.Password,
    }
    userRepo.Save(&user)
    return &dto.AuthenticationDTO{
        Token: strconv.Itoa(user.UserId.Id),
    }
}

// PasswordLogin 账号密码登录
func (userApplicationService UserApplicationService) PasswordLogin(command *command.PasswordLoginCommand) *dto.AuthenticationDTO {
    var user = userApplicationService.userRepository.FindByUsername(command.Username)
    if user.Password != sha.Encrypt(command.Password) {
        panic("密码不正确")
    }
    return &dto.AuthenticationDTO{
        Token: strconv.Itoa(user.UserId.Id),
    }
}
