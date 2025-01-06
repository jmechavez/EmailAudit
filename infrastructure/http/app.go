package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jmechavez/EmailAudit/infrastructure/db"
	"github.com/jmechavez/EmailAudit/internal/ports/service"
)

func Start() {
	router := mux.NewRouter()

	uh := UserHandlers{
		service.NewUserService(db.NewUserRepoStub()),
	}
	router.HandleFunc("/users", uh.GetAllUser).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
