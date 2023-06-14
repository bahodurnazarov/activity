package main

import (
	"github.com/bahodurnazarov/activity/internal/handler"
	lg "github.com/bahodurnazarov/activity/pkg/utils"
)

// @title Todo App API
// @version 1.0
// @description API Server for TodoList Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	lg.Errl.Println("Hello")
	handler.Run()
}
