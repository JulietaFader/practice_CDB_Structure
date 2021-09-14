package reader

import (
	"github.com/mercadolibre/fury_asset-mgmt-core-cdb/pkg/rest"
	"github.com/mercadolibre/fury_asset-mgmt-core-libs/pkg/transport/httpcore"
	"github.com/mercadolibre/fury_go-platform/pkg/fury"
)

func (h *handler) RouteURLs(app *fury.Application) {
	group := app.Router.Group("/api/crudnew/users")

	group.Get("/{id}", h.GetByID, httpcore.Middle(app))
	group.Put("/{id}", h.UpdateUser, httpcore.Middle(app))
	group.Delete("/{id}", h.DeleteUser, httpcore.Middle(app))
	group.Get("/", h.GetUsers, httpcore.Middle(app))
	group.Post("/", h.CreateUser, httpcore.Middle(app))
}

func (h *handler) API() *fury.Application {
	return rest.API(h)
}
