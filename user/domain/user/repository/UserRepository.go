package repository

import "github.com/changqing98/cqzone/user/domain/user/model"

type UserRepository interface {
    NextUserId() string
    Save(user model.User)
    FindByUserId(userId string) model.User
    FindByMobile(mobile string) model.User
    FindByEmail(email string) model.User
}
