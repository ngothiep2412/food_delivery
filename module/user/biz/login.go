package userbiz

import (
	"context"
	"g05-food-delivery/common"
	tokenprovider "g05-food-delivery/component/tokenprovider"
	usermodel "g05-food-delivery/module/user/model"
)

type LoginStorage interface {
	FindUser(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*usermodel.User, error)
}

type LoginBiz struct {
	store         LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBiz(store LoginStorage,
	tokenProvider tokenprovider.Provider,
	hasher Hasher,
	expiry int,
) *LoginBiz {
	return &LoginBiz{
		store:         store,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (biz *LoginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	passHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password == passHashed {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
