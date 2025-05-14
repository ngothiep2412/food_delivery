package appctx

import (
	"g05-food-delivery/component/uploadprovider"
	"g05-food-delivery/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMaiDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubSub() pubsub.Pubsub
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
	ps             pubsub.Pubsub
}

func NewAppContext(db *gorm.DB,
	provider uploadprovider.UploadProvider,
	secretKey string,
	ps pubsub.Pubsub) *appCtx {
	return &appCtx{db: db,
		uploadProvider: provider,
		secretKey:      secretKey,
		ps:             ps,
	}
}

func (c *appCtx) GetMaiDBConnection() *gorm.DB {
	return c.db
}
func (c *appCtx) UploadProvider() uploadprovider.UploadProvider { return c.uploadProvider }
func (c *appCtx) SecretKey() string {
	return c.secretKey
}
func (c *appCtx) GetPubSub() pubsub.Pubsub {
	return c.ps
}
