package reader

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mercadolibre/fury_go-core/pkg/web"

	"github.com/mercadolibre/crudnew/pkg/domain/process"
	"github.com/mercadolibre/crudnew/pkg/rest"
)

type handler struct {
	serv process.Service
}

func NewHandler(serv1 process.Service) rest.Handler {
	return &handler{serv: serv1}
}

type Handler interface {
	GetUsers(w http.ResponseWriter, req *http.Request) error
	GetByID(w http.ResponseWriter, req *http.Request) error
	CreateUser(w http.ResponseWriter, req *http.Request) error
	UpdateUser(w http.ResponseWriter, req *http.Request) error
	DeleteUser(w http.ResponseWriter, req *http.Request) error
}

func (h *handler) GetUsers(w http.ResponseWriter, req *http.Request) error {
	ctx := req.Context()
	users, err := h.serv.GetUsers(ctx)

	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
		return err
	}

	return web.RespondJSON(w, users, http.StatusOK)
}

func (h *handler) GetByID(w http.ResponseWriter, req *http.Request) error {
	ctx := req.Context()
	id := web.Params(req)["id"]

	user, err := h.serv.GetByID(ctx, id)

	if err != nil {
		_, _ = fmt.Fprint(w, err.Error())
		return err
	}

	return web.RespondJSON(w, user, http.StatusOK)
}

func (h *handler) CreateUser(w http.ResponseWriter, req *http.Request) error {
	ctx := req.Context()

	user := &process.User{}
	err := json.NewDecoder(req.Body).Decode(user)

	if err != nil {
		_, _ = fmt.Fprintf(w, "Can not parser")
		return err
	}

	msj, err := h.serv.CreateUser(ctx, user)

	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
		return err
	}

	_, _ = fmt.Fprintf(w, msj)
	return web.RespondJSON(w, msj, http.StatusOK)
}

func (h *handler) UpdateUser(w http.ResponseWriter, req *http.Request) error {
	ctx := req.Context()
	_user:= &process.User{}
	id := web.Params(req)["id"]

	err := json.NewDecoder(req.Body).Decode(_user)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Can not parser")
		return err
	}

	user, err := h.serv.UpdateUser(ctx, id, _user)

	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
		return err
	}

	return web.RespondJSON(w, user, http.StatusOK)
}

func (h *handler) DeleteUser(w http.ResponseWriter, req *http.Request) error {
	ctx := req.Context()
	id := web.Params(req)["id"]

	msj, err := h.serv.DeleteUser(ctx, id)

	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
		return err
	}

	_, _ = fmt.Fprintf(w, msj)
	return web.RespondJSON(w, msj, http.StatusOK)
}
