package command

type PasswordLoginCommand struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
