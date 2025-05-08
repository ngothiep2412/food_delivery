package userstorage

import (
	"context"
	"errors"
	"g05-food-delivery/common"
	usermodel "g05-food-delivery/module/user/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(ctx context.Context,
	conditions map[string]interface{},
	moreInfo ...string,
) (*usermodel.User, error) {
	userDB := usermodel.User{}
	db := s.db.Table(userDB.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
	}

	return &user, nil
}
