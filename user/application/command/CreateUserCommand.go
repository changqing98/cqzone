package command

type EmailRegisterCommand struct {
    Email            string `json:"email"`
    VerificationCode string `json:"verificationCode"`
    Password         string `json:"password"`
}
