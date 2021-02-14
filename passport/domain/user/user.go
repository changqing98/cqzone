package user

type User struct {
	UserId   *UserId
	Nickname string
	Password string
	Email    string
}

type UserId struct {
    Id int64
}
