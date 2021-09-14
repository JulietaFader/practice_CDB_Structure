package rest

import (
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/transport/httpcore"
	"github.com/mercadolibre/fury_go-core/pkg/log"
	"github.com/mercadolibre/fury_go-platform/pkg/fury"
)

type Handler interface {
	API() *fury.Application
	RouteURLs(app *fury.Application)
}

func API(handler Handler) *fury.Application {

	app, err := fury.NewWebApplication(
		fury.WithLogOptions(
			log.StacktraceOnError(false),
			log.WithCaller(false),
		),
		fury.WithErrorHandler(httpcore.ErrorHandler),
		fury.WithErrorEncoder(httpcore.ErrorEncoder),
	)
	if err != nil {
		panic(err.Error())
	}

	handler.RouteURLs(app)
	return app
}
