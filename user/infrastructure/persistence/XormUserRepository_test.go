package persistence

import (
    "github.com/changqing98/cqzone/user/domain/user/model"
    "github.com/changqing98/cqzone/user/infrastructure/do"
    StringUtils "github.com/changqing98/cqzone/user/infrastructure/utils/string"
    "github.com/changqing98/cqzone/user/infrastructure/xorm"
    "testing"
)

func TestUserRepositoryImpl_save(t *testing.T) {
    var userDo do.UserDO
    engine := xorm.GetXORMEngine()
    engine.Sync(&userDo)
    userRepositoryImpl := UserRepositoryImpl{
        Engine: engine,
    }
    email := StringUtils.RandomString(6) + "@test.com"
    var user = &model.User{
        Mobile:   "1" + StringUtils.RandomIntString(10),
        Email:    email,
        Password: "123456",
        Nickname: "nickname",
    }
    userRepositoryImpl.Save(user)
    user = userRepositoryImpl.FindByUserId(user.UserId.Id)
}
