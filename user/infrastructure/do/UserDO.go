package do

import "time"

type UserDO struct {
    Id        int64
    UserId    int64     `xorm:"notnull"`
    Nickname  string    `xorm:"varchar(50) notnull"`
    Mobile    string    `xorm:"varchar(16) notnull unique(uni_user_mobile)"`
    Email     string    `xorm:"varchar(40) notnull unique(uni_user_email)"`
    Password  string    `xorm:"varchar(60) notnull"`
    CreatedAt time.Time `xorm:"created notnull"`
    UpdatedAt time.Time `xorm:"updated notnull"`
}

func (UserDO) TableName() string {
    return "user"
}
