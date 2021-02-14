package data

import "time"

type UserDO struct {
    Id        int
    UserId    int64     `xorm:"bigint notnull unique(uk_user_user_id)"`
    Nickname  string    `xorm:"varchar(20) notnull"`
    Password  string    `xorm:"varchar(60) notnull"`
    Email     string    `xorm:"varchar(40) notnull unique(uk_user_email)"`
    CreatedAt time.Time `xorm:"timestamp created notnull"`
    UpdatedAt time.Time `xorm:"timestamp updated notnull"`
}

func (UserDO) TableName() string {
    return "user"
}
