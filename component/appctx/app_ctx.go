package appctx

import (
	"g05-food-delivery/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMaiDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, provider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{db: db, uploadProvider: provider}
}

func (c *appCtx) GetMaiDBConnection() *gorm.DB {
	return c.db
}
func (c *appCtx) UploadProvider() uploadprovider.UploadProvider { return c.uploadProvider }
