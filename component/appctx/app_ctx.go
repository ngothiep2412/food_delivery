package appctx

import "gorm.io/gorm"

type AppContext interface {
	GetMaiDBConnection() *gorm.DB
}

type appCtx struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *appCtx {
	return &appCtx{db: db}
}

func (c *appCtx) GetMaiDBConnection() *gorm.DB {
	return c.db
}
