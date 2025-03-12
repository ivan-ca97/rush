package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/ivan-ca97/rush/backend/controllers"
	"github.com/ivan-ca97/rush/backend/middlewares"
	"github.com/ivan-ca97/rush/backend/repositories"
	rush_db "github.com/ivan-ca97/rush/backend/repositories/db"
	"github.com/ivan-ca97/rush/backend/services"
)

func initLog(logDirectory string) *os.File {
	timestamp := time.Now().Format("20060102_150405")
	logFileName := fmt.Sprintf("Rush_%s.log", timestamp)

	logFilePath := logDirectory + "/" + logFileName

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}

	log.SetOutput(logFile)
	return logFile
}

func RushServer(port int, logDirectory string, jwtKey string, tokenExpirationTime time.Duration) {
	logFile := initLog(logDirectory)
	defer logFile.Close()

	auth := middlewares.AuthenticationFactory(jwtKey, tokenExpirationTime)

	db := rush_db.SetupDb()
	repositories := repositories.RushRepositories{Db: db}
	services := services.RushServices{Repositories: &repositories}
	controllers := controllers.RushControllers{Services: &services}

	r := chi.NewMux()

	r.Group(func(r chi.Router) {
		r.Use(auth.AuthenticationContext)
		r.Post("/authentication/login", controllers.Login)
		r.Post("/authentication/register", controllers.Register)
	})

	r.Group(func(r chi.Router) {
		r.Use(middlewares.LoggingMiddleware)
		r.Use(auth.AuthenticationMiddleware)
		r.Use(auth.AuthenticationContext)
		r.Get("/authentication/check", check)
		r.Get("/heartbeat", getHeartbeat)
		r.Get("/members", getHeartbeat)
	})

	address := fmt.Sprintf(":%d", port)
	http.ListenAndServe(address, r)
}

func getHeartbeat(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("beating"))
}

func check(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("wenas"))
}
