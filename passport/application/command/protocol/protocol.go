package protocol

type CreateUserCommand struct {
    // 邮箱
    Email string
    // 验证码
    Code int
}

type PasswordLoginCommand struct {
    // 邮箱
    Email string
    // 密码
    Password string
}

type AuthorizeCommand struct {
    ResponseType string
    Scope        string
    ClientId     string
    RedirectURI  string
    State        string
}

type FetchTokenCommand struct {
    GrantType string
    // The authorization code received from the authorization server.
    Code string
}

type AuthenticationDTO struct {
    RefreshToken string `json:"refresh_token"`
    AccessToken string `json:"access_token"`
}
