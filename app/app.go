package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jmechavez/EmailAudit/domain"
	"github.com/jmechavez/EmailAudit/service"
)

func Start() {
	router := mux.NewRouter()

	uh := UserHandlers{
		service.NewUserService(domain.NewUserRepoStub()),
	}
	router.HandleFunc("/users", uh.GetAllUser).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
