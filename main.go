package main

import (
	"fmt"
	"time"

	"github.com/JoaumVictor/pentakill-backend/config"
	"github.com/JoaumVictor/pentakill-backend/internal/models"
)


func main() {
	fmt.Println("Iniciando a aplicação...")

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Game{})

	fmt.Println("Conexão com o banco de dados foi bem-sucedida!")

	for {
		time.Sleep(time.Hour)
	}
}
