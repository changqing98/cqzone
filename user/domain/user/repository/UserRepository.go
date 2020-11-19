package repository

import "github.com/changqing98/cqzone/user/domain/user/model"

type UserRepository interface {
    NextUserId() int64
    Save(user *model.User)
    FindByUserId(userId int64) *model.User
    FindByMobile(mobile string) *model.User
    FindByEmail(email string) *model.User
}
