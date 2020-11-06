package service

import (
    "github.com/changqing98/cqzone/user/infrastructure/utils/logger"
)

// SmsServiceImpl 短信服务实现
type SmsServiceImpl struct {
}

var smsService SmsService

// CreateSmsService 创建SmsService
func CreateSmsService() SmsService {
    if smsService == nil {
        smsService = &SmsServiceImpl{}
    }
    return smsService
}

// SendVerificationCode 发送短信验证码
func (smsService SmsServiceImpl) SendVerificationCode(mobile string, verificationCode string) {
    logger.Info("mobile: " + mobile + ", sms verification code: " + verificationCode)
}
