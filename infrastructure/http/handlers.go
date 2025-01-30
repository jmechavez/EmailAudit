package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/jmechavez/EmailAudit/internal/dto"
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

func (uh *UserHandlers) NewUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	emailIdStr := vars["email_id"]

	// Convert emailIdStr to int64
	emailId, err := strconv.ParseInt(emailIdStr, 10, 64)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, "Invalid email ID format")
		return
	}

	var req dto.NewUserRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		req.EmailId = emailId
		user, appError := uh.service.NewUser(req)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, user)
		}
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
