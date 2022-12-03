package app

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jimmmisss/banking-auth/domain"
	"github.com/jimmmisss/banking-auth/service"
	"github.com/jimmmisss/banking-auth/util"
	"github.com/jimmmisss/banking-lib/logger"
	"log"
	"net/http"
)

func Start() {
	env := util.EnvCheck()
	router := mux.NewRouter()

	authRepository := domain.NewAuthRepository(util.EnvDB())
	ah := AuthHandler{service.NewLoginService(authRepository, domain.GetRolePermissions())}

	router.HandleFunc("/auth/login", ah.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/register", ah.NotImplementedHandler).Methods(http.MethodPost)
	router.HandleFunc("/auth/refresh", ah.Refresh).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", ah.Verify).Methods(http.MethodGet)

	address := env.Server.Address
	port := env.Server.Port
	logger.Info(fmt.Sprintf("Start server on %s:%s", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", env.Server.Address, env.Server.Port), router))
}
