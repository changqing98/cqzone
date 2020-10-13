package command

type PasswordLoginCommand struct {
	Mobile string `json:"mobile"`
	Password string `json:"password"`
}