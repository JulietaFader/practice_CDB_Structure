package reader

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mercadolibre/crudnew/pkg/domain/process"

	//"github.com/mercadolibre/fury_go-core/pkg/web"
)

type handler struct {
	serv process.Service
}

func NewHandler(serv1 process.Service) Handler {

	return &handler{serv: serv1}
}

type Handler interface {
	GetUsers(w http.ResponseWriter, req *http.Request)
	GetByID(w http.ResponseWriter, req *http.Request)
	CreateUser(w http.ResponseWriter, req *http.Request)
	UpdateUser(w http.ResponseWriter, req *http.Request)
	DeleteUser(w http.ResponseWriter, req *http.Request)
}

func (h *handler) GetUsers(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	ctx := req.Context()

	users, err := h.serv.GetUsers(ctx)

	if err != nil {

		w.WriteHeader(400)
		fmt.Fprintf(w, err.Error())
	} else {

		json.NewEncoder(w).Encode(users)
		w.WriteHeader(200)
	}

}
func (h *handler) GetByID(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := req.Context()
	id := mux.Vars(req)["id"]
	user, err := h.serv.GetByID(ctx, id)

	if err != nil {
		w.WriteHeader(404)
		fmt.Fprint(w, err.Error())

	} else {
		json.NewEncoder(w).Encode(user)
		w.WriteHeader(200)

	}
}
func (h *handler) CreateUser(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	ctx := req.Context()
	var user process.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Can not parser")
	}
	mnsj, err := h.serv.CreateUser(ctx, &user)

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, err.Error())
	} else {

		w.WriteHeader(201)
		fmt.Fprintf(w, mnsj)

	}
}
func (h *handler) UpdateUser(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	ctx := req.Context()
	var user process.User
	id := mux.Vars(req)["id"]
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Can not parser")
	}

	response, err := h.serv.UpdateUser(ctx, id, &user)

	if err != nil {

		w.WriteHeader(404)
		fmt.Fprintf(w, err.Error())

	} else {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(response)
	}

}

func (h *handler) DeleteUser(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	ctx := req.Context()
	id := mux.Vars(req)["id"]

	msj, err := h.serv.DeleteUser(ctx, id)

	if err != nil {

		w.WriteHeader(404)
		fmt.Fprintf(w, err.Error())
	} else {
		w.WriteHeader(204)
		fmt.Fprintf(w, msj)
	}

}
