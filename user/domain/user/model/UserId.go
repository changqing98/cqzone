package model

type UserId struct {
    Id int
}

func NewUserId(Id int) *UserId {
    return &UserId{Id}
}

func (userId UserId) getId() int {
    return userId.Id
}
