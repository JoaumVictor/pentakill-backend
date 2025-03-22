package repository

import (
	"github.com/JoaumVictor/pentakill-backend/config"

	"github.com/JoaumVictor/pentakill-backend/internal/models"
)

func GetAllGames() ([]models.Game, error) {
	var games []models.Game
	result := config.DB.Find(&games)
	return games, result.Error
}

func CreateGame(game *models.Game) error {
	result := config.DB.Create(game)
	return result.Error
}

func UpdateGame(id uint, updatedGame models.Game) error {
	var game models.Game
	if err := config.DB.First(&game, id).Error; err != nil {
		return err
	}
	game.Name = updatedGame.Name
	game.Description = updatedGame.Description
	game.Genre = updatedGame.Genre
	game.Platform = updatedGame.Platform
	return config.DB.Save(&game).Error
}

func DeleteGame(id uint) error {
	return config.DB.Delete(&models.Game{}, id).Error
}
