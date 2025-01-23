package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jmechavez/EmailAudit/internal/ports/service"
)

type UserHandlers struct {
	service service.UserService
}

func (uh *UserHandlers) GetAllUser(w http.ResponseWriter, r *http.Request) {
	// status := r.URL.Query().Get()
	status := r.URL.Query().Get("status")

	users, err := uh.service.GetAllUser(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, users)
	}
}

func (uh *UserHandlers) GetUserNo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email_id := vars["email_id"]

	users, err := uh.service.ByUserNum(email_id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, users)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(fmt.Sprintf("failed to encode response: %v", err))
	}
}
