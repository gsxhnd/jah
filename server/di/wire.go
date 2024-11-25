//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package di

import (
	"github.com/google/wire"
	"github.com/gsxhnd/jaha/server/db"
	"github.com/gsxhnd/jaha/server/handler"
	"github.com/gsxhnd/jaha/server/middleware"
	"github.com/gsxhnd/jaha/server/router"
	"github.com/gsxhnd/jaha/server/service"
	"github.com/gsxhnd/jaha/server/storage"
	"github.com/gsxhnd/jaha/utils"
)

func InitApp() (*Application, error) {
	wire.Build(
		utils.UtilsSet,
		NewApplication,
		router.NewRouter,
		middleware.NewMiddleware,
		handler.HandlerSet,
		service.ServiceSet,
		storage.StorageSet,
		db.DBSet,
	)
	return &Application{}, nil
}
