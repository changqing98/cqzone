package service

import (
    "github.com/changqing98/cqzone/user/application/command"
    "github.com/changqing98/cqzone/user/application/dto"
    "github.com/changqing98/cqzone/user/domain/user/model"
    "github.com/changqing98/cqzone/user/domain/user/repository"
    infraSvc "github.com/changqing98/cqzone/user/infrastructure/service"
)

// userApplicationServiceImpl 用户应用服务
type userApplicationServiceImpl struct {
    userRepository repository.UserRepository
    smsService     infraSvc.SmsService
}

func NewUserApplicationService(userRepository repository.UserRepository, smsService infraSvc.SmsService) UserApplicationService {
    return userApplicationServiceImpl{userRepository, smsService}
}

// SendSmsVerificationCode 发送手机验证码的实现
func (userApplicationService userApplicationServiceImpl) SendSmsVerificationCode(sendSmsVerificationCodeCommand *command.SendSmsVerificationCodeCommand) string {
    smsVerificationCode := generateSmsVerificationCode()
    mobile := sendSmsVerificationCodeCommand.Mobile
    userApplicationService.smsService.SendVerificationCode(mobile, smsVerificationCode)
    return smsVerificationCode
}

// Register 用户注册
func (userApplicationService userApplicationServiceImpl) Register(createUserCommand *command.EmailRegisterCommand) *dto.AuthenticationDTO {
    userRepo := userApplicationService.userRepository
    user := model.User{
        Email:    createUserCommand.Email,
        Password: createUserCommand.Password,
    }
    userRepo.Save(&user)
    return &dto.AuthenticationDTO{
        Token: "token",
    }
}

// PasswordLogin 账号密码登录
func (userApplicationService userApplicationServiceImpl) PasswordLogin(passwordLoginCommand *command.PasswordLoginCommand) *dto.AuthenticationDTO {
    return &dto.AuthenticationDTO{}
}

// 生成验证码
func generateSmsVerificationCode() string {
    return "123456"
}
