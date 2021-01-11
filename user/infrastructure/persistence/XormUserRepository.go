package persistence

import (
    "github.com/changqing98/cqzone/user/domain/user/model"
    "github.com/changqing98/cqzone/user/infrastructure/data"
    _ "github.com/changqing98/cqzone/user/infrastructure/utils/snowflake"
    "github.com/go-xorm/xorm"
    "math/rand"
    "time"
)

type UserRepositoryImpl struct {
    Engine *xorm.Engine
}

func (userRepoImpl UserRepositoryImpl) NextUserId() *model.UserId {
    rand.Seed(time.Now().UnixNano())
    randNum := rand.Intn(111111)
    userIdVal := 10000000 + randNum
    return model.NewUserId(userIdVal)
}

func (userRepoImpl UserRepositoryImpl) Save(user *model.User) {
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

func (userRepoImpl UserRepositoryImpl) FindByUserId(userId *model.UserId) *model.User {
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

func (userRepoImpl UserRepositoryImpl) FindByUsername(username string) *model.User {
    var userDo data.UserDO
    success, err := userRepoImpl.Engine.Where("username = ?", username).Get(&userDo)
    if !success {
        panic("User not found")
    }
    if err != nil {
        panic(err)
    }
    return ConvertDOToUser(&userDo)
}

func ConvertUserToDo(user *model.User) *data.UserDO {
    var userDo = data.UserDO{
        UserId:   user.UserId.Id,
        Nickname: user.Nickname,
        Password: user.Password,
    }
    return &userDo
}

func ConvertDOToUser(userDO *data.UserDO) *model.User {
    userId := userDO.UserId
    var user = model.User{
        UserId:   &model.UserId{Id: userId},
        Nickname: userDO.Nickname,
        Password: userDO.Password,
    }
    return &user
}
