package repository

import "github.com/changqing98/cqzone/user/domain/user/model"

type UserRepository interface {
    save(user model.User) model.User
    findByUserId(userId string) model.User
    findByMobile(mobile string) model.User
    findByEmail(email string) model.User
}
