package model

type UserId struct {
    Id string
}

func NewUserId(id string) UserId {
    return UserId{id}
}
