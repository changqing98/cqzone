package service

// SmsService 短信服接口定义
type SmsService interface {
	// 发送手机验证码
	SendVerificationCode(mobile string, verificationCode string)
}


