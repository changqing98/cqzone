package persistence

import (
    "github.com/changqing98/cqzone/user/domain/user/model"
    "github.com/changqing98/cqzone/user/infrastructure/config"
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestUserRepositoryImpl_save(t *testing.T) {
    engine := config.GetXORMEngine()
    userRepositoryImpl := UserRepositoryImpl{
        Engine: engine,
    }
    userId := 123456
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
