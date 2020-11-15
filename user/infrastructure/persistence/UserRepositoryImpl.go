package persistence

import (
    "github.com/changqing98/cqzone/user/domain/user/model"
    "github.com/changqing98/cqzone/user/infrastructure/convertor"
    "github.com/changqing98/cqzone/user/infrastructure/do"
    "github.com/go-xorm/xorm"
)

type UserRepositoryImpl struct {
    Engine *xorm.Engine
}

func (userRepoImpl UserRepositoryImpl) NextUserId() string {
    return "user_id"
}

func (userRepoImpl UserRepositoryImpl) Save(user model.User) {
    userDo := convertor.ConvertUserToDo(user)
    _, err := userRepoImpl.Engine.Insert(userDo)
    if err != nil {
        panic(err)
    }
}

func (userRepoImpl UserRepositoryImpl) FindByUserId(userId string) model.User {
    var userDo do.UserDO
    success, err := userRepoImpl.Engine.Where("user_id = ?", userId).Get(&userDo)
    if !success {
        panic("User not found")
    }
    if err != nil {
        panic(err)
    }
    return convertor.ConvertDOToUser(userDo)
}

func (userRepoImpl UserRepositoryImpl) FindByMobile(mobile string) model.User {
    var userDo do.UserDO
    userRepoImpl.Engine.Where("mobile = ?", mobile).Get(&userDo)
    return convertor.ConvertDOToUser(userDo)
}

func (userRepoImpl UserRepositoryImpl) FindByEmail(email string) model.User {
    var userDo do.UserDO
    userRepoImpl.Engine.Where("email = ?", email).Get(&userDo)
    return convertor.ConvertDOToUser(userDo)
}
