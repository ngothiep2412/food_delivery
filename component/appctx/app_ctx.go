package appctx

import (
	"g05-food-delivery/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMaiDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
}

func NewAppContext(db *gorm.DB, provider uploadprovider.UploadProvider, secretKey string) *appCtx {
	return &appCtx{db: db, uploadProvider: provider, secretKey: secretKey}
}

func (c *appCtx) GetMaiDBConnection() *gorm.DB {
	return c.db
}
func (c *appCtx) UploadProvider() uploadprovider.UploadProvider { return c.uploadProvider }

func (c *appCtx) SecretKey() string {
	return c.secretKey
}
