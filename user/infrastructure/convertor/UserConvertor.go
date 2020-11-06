package convertor

import (
    "github.com/changqing98/cqzone/user/domain/user/model"
    "github.com/changqing98/cqzone/user/infrastructure/do"
)

func convert(user model.User) do.UserDO {
    var do = do.UserDO{
        UserId:             user.UserId.Id,
        Nickname:           user.Nickname,
        Mobile:             user.Mobile,
        Email:              user.Email,
        Password:           user.Password,
        HasEmailValidated:  user.HasEmailValidated,
        HasMobileValidated: user.HasEmailValidated,
    }
    return do
}
