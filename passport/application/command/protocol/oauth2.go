package protocol

type AuthorizeRequest struct {
    ResponseType string
    Scope        string
    ClientId     string
    RedirectURI  string
    State        string
}

type AuthorizeResponse struct {
    Code  string
    State string
}

type AccessTokenRequest struct {
    GrantType   string
    Code        string
    RedirectURI string
    ClientId    string
}

type AccessTokenResponse struct {
    RefreshToken string
    AccessToken  string
    TokenType    string
    ExpiresIn    uint32
    IdToken      string
}
