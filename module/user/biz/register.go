package userbiz

import (
	"context"
	"g05-food-delivery/common"
	usermodel "g05-food-delivery/module/user/model"
)

type RegisterStorage interface {
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
	FindUser(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)
}

type Hasher interface {
	Hash(data string) string
}

type registerBiz struct {
	store  RegisterStorage
	hasher Hasher
}

func NewRegisterBiz(registerStore RegisterStorage, hasher Hasher) *registerBiz {
	return &registerBiz{store: registerStore, hasher: hasher}
}

func (biz *registerBiz) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		//if user.Status == 0 {
		//	return error
		//}
		return usermodel.ErrEmailExists
	}

	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"
	//data.Status = 1

	if err := biz.store.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
