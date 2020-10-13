package service

import (
	"github.com/changqing98/cqzone/user/application/command"
	"github.com/changqing98/cqzone/user/application/dto"
	"github.com/changqing98/cqzone/user/infrastructure/service"
)

// UserApplicationService 用户serice
type IUserApplicationService interface {
	// 发送手机验证码
	SendSmsVerificationCode(sendSmsVerificationCodeCommand *command.SendSmsVerificationCodeCommand) string
	// 用户注册
	Register(createUserCommand *command.CreateUserCommand) dto.AuthenticationDto
	// 密码登录
	PasswordLogin(passwordLoginCommand *command.PasswordLoginCommand) dto.AuthenticationDto
}

// UserApplicationService 用户应用服务
type UserApplicationService struct {
	smsService *service.SmsService
}

var userApplicationService *UserApplicationService

// CreateUserApplicationService 创建单例UserApplicationService
func CreateUserApplicationService() *UserApplicationService {
	if userApplicationService == nil {
		userApplicationService = &UserApplicationService{
			smsService: service.CreateSmsService(),
		}
	}
	return userApplicationService
}

// SendSmsVerificationCode 发送手机验证码的实现
func (userApplicationService *UserApplicationService) SendSmsVerificationCode(sendSmsVerificationCodeCommand *command.SendSmsVerificationCodeCommand) string {
	smsVerificationCode := generateSmsVerificationCode()
	mobile := sendSmsVerificationCodeCommand.Mobile
	userApplicationService.smsService.SendVerificationCode(mobile, smsVerificationCode)
	return smsVerificationCode
}

func generateSmsVerificationCode() string {
	return "123456"
}
