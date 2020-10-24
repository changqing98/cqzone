package service

import (
	"user/application/command"
	"user/application/dto"
	"user/infrastructure/service"
)

// UserApplicationService 用户应用服务接口定义
type UserApplicationService interface {
	// 发送手机验证码
	SendSmsVerificationCode(sendSmsVerificationCodeCommand command.SendSmsVerificationCodeCommand) string
	// 用户注册
	Register(createUserCommand command.CreateUserCommand) dto.AuthenticationDto
	// 密码登录
	PasswordLogin(passwordLoginCommand command.PasswordLoginCommand) dto.AuthenticationDto
}

// UserApplicationServiceImpl 用户应用服务
type UserApplicationServiceImpl struct {
	smsService service.SmsService
}

var userApplicationService UserApplicationService

// CreateUserApplicationService 创建单例UserApplicationService
func CreateUserApplicationService() UserApplicationService {
	if userApplicationService == nil {
		userApplicationService = &UserApplicationServiceImpl{
			smsService: service.CreateSmsService(),
		}
	}
	return userApplicationService
}

// SendSmsVerificationCode 发送手机验证码的实现
func (userApplicationService UserApplicationServiceImpl) SendSmsVerificationCode(sendSmsVerificationCodeCommand command.SendSmsVerificationCodeCommand) string {
	smsVerificationCode := generateSmsVerificationCode()
	mobile := sendSmsVerificationCodeCommand.Mobile
	userApplicationService.smsService.SendVerificationCode(mobile, smsVerificationCode)
	return smsVerificationCode
}

// Register 用户注册
func (userApplicationService UserApplicationServiceImpl) Register(createUserCommand command.CreateUserCommand) dto.AuthenticationDto {
	return dto.AuthenticationDto{}
}

// PasswordLogin 账号密码登录
func (userApplicationService UserApplicationServiceImpl) PasswordLogin(passwordLoginCommand command.PasswordLoginCommand) dto.AuthenticationDto {
	return dto.AuthenticationDto{}
}

// 生成验证码
func generateSmsVerificationCode() string {
	return "123456"
}
