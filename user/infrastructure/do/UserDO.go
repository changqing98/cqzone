package do

import "time"

type UserDO struct {
    Id                 int32
    UserId             string
    Nickname           string
    Mobile             string
    Email              string
    Password           string
    CreatedAt          time.Time
    UpdatedAt          time.Time
}
