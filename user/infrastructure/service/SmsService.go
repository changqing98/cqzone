package service

import (
	"github.com/changqing98/cqzone/user/infrastructure/utils/log"
)

// ISmsService 短信服接口定义
type ISmsService interface {
	// 发送手机验证码
	SendVerificationCode(mobile string, verificationCode string)
}

// SmsService 短信服务
type SmsService struct {
}

var smsService *SmsService

func CreateSmsService() *SmsService {
	if smsService == nil {
		smsService = &SmsService{}
	}
	return smsService
}

// SendVerificationCode 发送短信验证码
func (smsService *SmsService) SendVerificationCode(mobile string, verificationCode string) {
	log.Info("mobile: " + mobile + ", sms verification code: " + verificationCode)
}
