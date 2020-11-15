package do

import "time"

type UserDO struct {
    Id        int
    UserId    int
    Nickname  string
    Mobile    string
    Email     string
    Password  string
    CreatedAt time.Time `xorm:"created_at"`
    UpdatedAt time.Time `xorm:"updated_at"`
}

func (UserDO) TableName() string {
    return "user"
}
