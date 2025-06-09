package middleware

import (
	"errors"
	"fmt"
	"g05-food-delivery/common"
	"g05-food-delivery/component/appctx"
	"g05-food-delivery/component/tokenprovider/jwt"
	usermodel "g05-food-delivery/module/user/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"strings"
)

type AuthenStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(err,
		fmt.Sprintf("wrong authen header"),
		"ErrWrongAuthHeader",
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

func RequireAuth(ctx appctx.AppContext, authenStore AuthenStore) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJwtProvider(ctx.SecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.Request.Header.Get("Authorization"))

		if err != nil {
			panic(err)
		}

		//db := ctx.GetMaiDBConnection()
		//store := userstorage.NewSQLStore(db)
		//
		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}
		//
		//user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		user, err := authenStore.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		user.Mask(false)

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
