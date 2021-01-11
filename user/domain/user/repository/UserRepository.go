package repository

import "github.com/changqing98/cqzone/user/domain/user/model"

type UserRepository interface {
    NextUserId() *model.UserId
    Save(user *model.User)
    FindByUserId(userId *model.UserId) *model.User
    FindByUsername(username string) *model.User
}
