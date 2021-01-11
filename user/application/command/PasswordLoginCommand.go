package command

type PasswordLoginCommand struct {
    // 用户名
	Username string `json:"username"`
	// 密码
	Password string `json:"password"`
}
