package server

import (
	"kratos-admin/internal/conf"
	"kratos-admin/internal/interfaces"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	transHttp "github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server,
	user *interfaces.UserUseCase,
	logger log.Logger) *transHttp.Server {
	var opts = []transHttp.ServerOption{
		transHttp.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			metrics.Server(),
			validate.Validator(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, transHttp.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, transHttp.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, transHttp.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := transHttp.NewServer(opts...)

	srv.HandlePrefix("/", interfaces.RegisterHTTPServer(user))

	return srv
}
