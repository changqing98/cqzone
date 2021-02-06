package persistence

import (
    user2 "github.com/changqing98/cqzone/user/domain/user"
    "github.com/changqing98/cqzone/user/infrastructure/data"
    "github.com/changqing98/cqzone/user/infrastructure/xorm"
    "github.com/go-playground/assert/v2"
    "testing"
)

func TestUserRepositoryImpl_save(t *testing.T) {
    var userDo data.UserDO
    engine := xorm.GetXormEngine()
    engine.Sync(&userDo)
    userRepositoryImpl := UserRepositoryImpl{
        Engine: engine,
    }
    var nickname = "nickname"
    var password = "abcdefg"
    var user = &user2.User{
        Password: password,
        Nickname: nickname,
    }
    userRepositoryImpl.Save(user)
    user = userRepositoryImpl.FindByUserId(user.UserId)
    assert.Equal(t, user.Nickname, nickname)
    assert.Equal(t, user.Password, password)
}
