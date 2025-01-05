package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/jmechavez/EmailAudit/service"
)

type UserHandlers struct {
	service service.UserService
}

func (uh *UserHandlers) GetAllUser(w http.ResponseWriter, r *http.Request) {
	// status := r.URL.Query().Get()

	user, _ := uh.service.GetAllUser()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(user)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
