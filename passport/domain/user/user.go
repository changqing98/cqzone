package user

type User struct {
	UserId   UserId
	Username string
	Nickname string
	Password string
	Email    string
	Mobile   string
}

type UserId struct {
    Id int64
}
