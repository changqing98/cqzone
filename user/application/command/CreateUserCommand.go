package command

type CreateUserCommand struct {
	Mobile           string
	VerificationCode string
	password         string
}
