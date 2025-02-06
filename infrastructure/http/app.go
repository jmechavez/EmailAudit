package http

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/jmechavez/EmailAudit/infrastructure/db"
	"github.com/jmechavez/EmailAudit/internal/ports/service"
)

func Start() {
	router := mux.NewRouter()

	// uh := UserHandlers{
	// 	service.NewUserService(db.NewUserRepoStub()),
	// }
	dbUser := getDbUser()
	userRepositoryDb := db.NewUserRepoDb(dbUser)
	uh := UserHandlers{
		service.NewUserService(userRepositoryDb),
	}

	router.HandleFunc("/users", uh.GetAllUser).Methods(http.MethodGet)
	router.HandleFunc("/users/{email_id}", uh.GetUserNo).Methods(http.MethodGet)
	router.HandleFunc("/users/{id_no}/postuser", uh.NewUser).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getDbUser() *sqlx.DB {
	connStr := "user=admin password=admin123 dbname=email_dir sslmode=disable"
	userDb, err := sqlx.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return userDb
}
