package user

type UserRepository interface {
    NextUserId() *UserId
    Save(user *User)
    FindByUserId(userId *UserId) *User
    FindByUsername(username string) *User
}
