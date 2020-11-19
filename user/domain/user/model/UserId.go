package model

type UserId struct {
    Id int64
}

func NewUserId(id int64) UserId {
    return UserId{id}
}
