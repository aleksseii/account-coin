package main

import (
	"fmt"
	"github.com/go-chi/chi"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {

	logger.SetReportCaller(true)
	router := chi.NewRouter()
	printWelcomeMsg()

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}

}

func printWelcomeMsg() {
	fmt.Println("Starting account-coin service...")
	banner, err := os.ReadFile("./resources/banner.txt")
	if err != nil {
		logger.Fatalf("Failed to read banner, %v", err)
	}
	fmt.Println(string(banner))
}
