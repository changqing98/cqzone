package persistence

import (
    "github.com/changqing98/cqzone/user/domain/user/model"
)

type UserRepositoryImpl struct{}

func (UserRepositoryImpl) save(user model.User) model.User {
    panic("implement me")
}

func (UserRepositoryImpl) findByUserId(userId string) model.User {
    panic("implement me")
}

func (UserRepositoryImpl) findByMobile(mobile string) model.User {
    panic("implement me")
}

func (UserRepositoryImpl) findByEmail(email string) model.User {
    panic("implement me")
}


