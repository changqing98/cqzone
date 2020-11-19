package convertor

import (
    "github.com/changqing98/cqzone/user/domain/user/model"
    "github.com/changqing98/cqzone/user/infrastructure/do"
)

func ConvertUserToDo(user *model.User) *do.UserDO {
    var userDo = do.UserDO{
        UserId:   user.UserId.Id,
        Nickname: user.Nickname,
        Mobile:   user.Mobile,
        Email:    user.Email,
        Password: user.Password,
    }
    return &userDo
}

func ConvertDOToUser(userDO *do.UserDO) *model.User {
    userId := userDO.UserId
    var user = model.User{
        UserId:   &model.UserId{Id: userId},
        Nickname: userDO.Nickname,
        Mobile:   userDO.Mobile,
        Email:    userDO.Email,
        Password: userDO.Password,
    }
    return &user
}
