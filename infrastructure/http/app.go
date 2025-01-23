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

	// uh := UserHandlers{
	// 	service.NewUserService(db.NewUserRepoStub()),
	// }
	uh := UserHandlers{
		service.NewUserService(db.NewUserRepoDb()),
	}

	router.HandleFunc("/users", uh.GetAllUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{email_id}", uh.GetUserNo).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
