package command

type SendSmsVerificationCodeCommand struct{
	Mobile string `json:"mobile"`
}