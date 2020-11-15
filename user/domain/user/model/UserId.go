package model

type UserId struct {
    Id int
}

func NewUserId(id int) UserId {
    return UserId{id}
}
