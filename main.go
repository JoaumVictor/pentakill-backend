package main

import (
	"fmt"
	"time"

	"github.com/JoaumVictor/pentakill-backend/config"
	"github.com/JoaumVictor/pentakill-backend/internal/models"
	"github.com/JoaumVictor/pentakill-backend/internal/routes"
)


func main() {
	fmt.Println("Iniciando a aplicação...")

	// conectei no banco
	config.ConnectDatabase()

	// migrei meu modelo games no banco
	config.DB.AutoMigrate(&models.Game{})
	err := config.DB.AutoMigrate(&models.Game{})
	if err != nil {
		fmt.Println("Erro ao rodar a migration:", err)
		return
	}

	fmt.Println("Conexão com o banco de dados foi bem-sucedida!")

	// se tudo deu certo, eu inicio meu backend
	r := routes.SetupRouter()
	r.Run(":8080")

	// se meu banco iniciar antes do backend, eu dou esse timesleep pra esperar um pouco
	for {
		time.Sleep(time.Hour)
	}
}
