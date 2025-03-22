package handlers

import (
	"net/http"
	"strconv"

	"github.com/JoaumVictor/pentakill-backend/internal/models"
	"github.com/JoaumVictor/pentakill-backend/internal/repository"

	"github.com/gin-gonic/gin"
)

func GetGames(c *gin.Context) {
	games, err := repository.GetAllGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar jogos"})
		return
	}
	c.JSON(http.StatusOK, games)
}

func CreateGame(c *gin.Context) {
	var game models.Game
	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	if err := repository.CreateGame(&game); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar jogo"})
		return
	}

	c.JSON(http.StatusCreated, game)
}

func UpdateGame(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedGame models.Game
	if err := c.ShouldBindJSON(&updatedGame); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	if err := repository.UpdateGame(uint(id), updatedGame); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar jogo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Jogo atualizado com sucesso"})
}

func DeleteGame(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := repository.DeleteGame(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar jogo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Jogo deletado com sucesso"})
}
