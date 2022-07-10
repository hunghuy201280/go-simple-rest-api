package component

import "gorm.io/gorm"

type AppContext interface {
	GetMainDbConnection() *gorm.DB
}

type appCtx struct {
	db *gorm.DB
}

func (ctx appCtx) GetMainDbConnection() *gorm.DB {
	return ctx.db
}

func NewAppContext(db *gorm.DB) *appCtx {
	return &appCtx{db: db}
}
