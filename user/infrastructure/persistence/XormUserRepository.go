package persistence

import (
    "github.com/changqing98/cqzone/user/domain/user/model"
    "github.com/changqing98/cqzone/user/infrastructure/convertor"
    "github.com/changqing98/cqzone/user/infrastructure/do"
    "github.com/changqing98/cqzone/user/infrastructure/utils/snowflake"
    "github.com/go-xorm/xorm"
)

type UserRepositoryImpl struct {
    Engine *xorm.Engine
}

func (userRepoImpl UserRepositoryImpl) NextUserId() int64 {
    return snowflake.Generate()
}

func (userRepoImpl UserRepositoryImpl) Save(user *model.User) {
    userId := user.UserId
    if userId == nil {
        user.UserId = &model.UserId{
            Id: userRepoImpl.NextUserId(),
        }
        userDo := convertor.ConvertUserToDo(user)
        _, err := userRepoImpl.Engine.Insert(userDo)
        if err != nil {
            panic(err)
        }
    } else {
        _, err := userRepoImpl.Engine.Where("user_id=?", userId.Id).Update(&do.UserDO{Mobile: user.Mobile, Password: user.Password, Nickname: user.Nickname, Email: user.Email})
        if err != nil {
            panic(err)
        }
    }

}

func (userRepoImpl UserRepositoryImpl) FindByUserId(userId int64) *model.User {
    var userDo do.UserDO
    success, err := userRepoImpl.Engine.Where("user_id = ?", userId).Get(&userDo)
    if !success {
        panic("User not found")
    }
    if err != nil {
        panic(err)
    }
    return convertor.ConvertDOToUser(&userDo)
}

func (userRepoImpl UserRepositoryImpl) FindByMobile(mobile string) *model.User {
    var userDo do.UserDO
    userRepoImpl.Engine.Where("mobile = ?", mobile).Get(&userDo)
    return convertor.ConvertDOToUser(&userDo)
}

func (userRepoImpl UserRepositoryImpl) FindByEmail(email string) *model.User {
    var userDo do.UserDO
    userRepoImpl.Engine.Where("email = ?", email).Get(&userDo)
    return convertor.ConvertDOToUser(&userDo)
}
