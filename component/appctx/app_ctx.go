package appctx

import (
	"g05-food-delivery/component/uploadprovider"
	"g05-food-delivery/pubsub"
	"g05-food-delivery/skio"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMaiDBConnection() *gorm.DB
	UploadProvider() uploadprovider.UploadProvider
	SecretKey() string
	GetPubSub() pubsub.Pubsub
	GetRealtimeEngine() skio.RealtimeEngine
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
	secretKey      string
	ps             pubsub.Pubsub
	rlEngine       skio.RealtimeEngine
}

func NewAppContext(db *gorm.DB,
	provider uploadprovider.UploadProvider,
	secretKey string,
	ps pubsub.Pubsub,
) *appCtx {
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

func (c *appCtx) GetRealtimeEngine() skio.RealtimeEngine {
	return c.rlEngine
}

func (c *appCtx) SetRealtimeEngine(rltEngine skio.RealtimeEngine) {
	c.rlEngine = rltEngine
}
