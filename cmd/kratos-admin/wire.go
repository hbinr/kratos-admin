// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"kratos-admin/internal/biz"
	"kratos-admin/internal/conf"
	"kratos-admin/internal/data"
	"kratos-admin/internal/server"
	"kratos-admin/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		newApp))
}
