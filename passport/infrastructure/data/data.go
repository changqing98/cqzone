package data

import "time"

type UserDO struct {
    Id        int
    UserId    int64     `xorm:"notnull unique(uk_user_user_id)"`
    Username  string    `xorm:"varchar(20) notnull"`
    Nickname  string    `xorm:"varchar(20) notnull"`
    Password  string    `xorm:"varchar(60) notnull"`
    CreatedAt time.Time `xorm:"created notnull"`
    UpdatedAt time.Time `xorm:"updated notnull"`
}

func (UserDO) TableName() string {
    return "passport"
}
