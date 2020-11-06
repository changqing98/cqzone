package do

import "time"

type UserDO struct {
    Id                 int32
    UserId             string
    Nickname           string
    Mobile             string
    Email              string
    Password           string
    HasMobileValidated bool
    HasEmailValidated  bool
    CreatedAt          time.Time
    UpdatedAt          time.Time
}
