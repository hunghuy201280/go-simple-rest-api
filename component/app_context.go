package component

import (
	"gorm.io/gorm"
	"simple-rest-api/component/uploadprovider"
)

type AppContext interface {
	GetMainDbConnection() *gorm.DB
	GetUploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider uploadprovider.UploadProvider
}

func (ctx appCtx) GetMainDbConnection() *gorm.DB {
	return ctx.db
}

func (ctx appCtx) GetUploadProvider() uploadprovider.UploadProvider {
	return ctx.uploadProvider
}

func NewAppContext(
	db *gorm.DB,
	uploadProvider uploadprovider.UploadProvider,
) *appCtx {
	return &appCtx{db: db, uploadProvider: uploadProvider}
}
