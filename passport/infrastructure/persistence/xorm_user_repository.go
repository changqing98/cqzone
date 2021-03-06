package persistence

import (
    "github.com/changqing98/cqzone/user/domain/user"
    "github.com/changqing98/cqzone/user/infrastructure/constant"
    "github.com/changqing98/cqzone/user/infrastructure/data"
    "github.com/changqing98/cqzone/user/infrastructure/utils/snowflake"
    _ "github.com/changqing98/cqzone/user/infrastructure/utils/snowflake"
    "github.com/go-xorm/xorm"
)

type UserRepositoryImpl struct {
    Engine *xorm.Engine
}

func (userRepoImpl UserRepositoryImpl) NextUserId() *user.UserId {
    return &user.UserId{
        Id: snowflake.Generate(),
    }
}

func (userRepoImpl UserRepositoryImpl) Save(user *user.User) {
    userId := user.UserId
    if userId == nil {
        user.UserId = userRepoImpl.NextUserId()
        userDo := ConvertUserToDo(user)
        _, err := userRepoImpl.Engine.Insert(userDo)
        if err != nil {
            panic(err)
        }
    } else {
        _, err := userRepoImpl.Engine.Where("user_id=?", userId.Id).Update(&data.UserDO{Password: user.Password, Nickname: user.Nickname})
        if err != nil {
            panic(err)
        }
    }
}

func (userRepoImpl UserRepositoryImpl) FindByUserId(userId *user.UserId) *user.User {
    var userDo data.UserDO
    success, err := userRepoImpl.Engine.Where("user_id = ?", userId.Id).Get(&userDo)
    if !success {
        panic("User not found")
    }
    if err != nil {
        panic(err)
    }
    return ConvertDOToUser(&userDo)
}

func (userRepoImpl UserRepositoryImpl) FindByEmail(email string) *user.User {
    var userDo data.UserDO
    success, err := userRepoImpl.Engine.Where("email = ?", email).Get(&userDo)
    if !success {
        panic(constant.UserNotFound)
    }
    if err != nil {
        panic(err)
    }
    return ConvertDOToUser(&userDo)
}

func ConvertUserToDo(user *user.User) *data.UserDO {
    var userDo = data.UserDO{
        UserId:   user.UserId.Id,
        Nickname: user.Nickname,
        Email: user.Email,
        Password: user.Password,
    }
    return &userDo
}

func ConvertDOToUser(userDO *data.UserDO) *user.User {
    userId := userDO.UserId
    return &user.User{
        UserId:   &user.UserId{Id: userId},
        Nickname: userDO.Nickname,
        Password: userDO.Password,
    }
}
