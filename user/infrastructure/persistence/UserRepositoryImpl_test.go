package persistence

import (
    "github.com/changqing98/cqzone/user/domain/user/model"
    "github.com/changqing98/cqzone/user/infrastructure/config"
    "github.com/changqing98/cqzone/user/infrastructure/do"
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestUserRepositoryImpl_save(t *testing.T) {
    var userDO do.UserDO
    engine := config.GetXORMEngine()
    engine.Sync2(&userDO)
    userRepositoryImpl := UserRepositoryImpl{
        Engine: engine,
    }
    userId := "test_user_id"
    email := "test@test.com"
    var user = model.User{
        UserId:   model.UserId{userId},
        Mobile:   "12345678910",
        Email:    email,
        Password: "123456",
        Nickname: "nickname",
    }
    userRepositoryImpl.Save(user)
    user = userRepositoryImpl.FindByUserId(email)
    assert.Equal(t, userId, user.UserId.Id)
}
