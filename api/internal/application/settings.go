package application

import (
	"github.com/mark-by/little-busy-back/api/internal/domain/entity"
	"github.com/mark-by/little-busy-back/api/internal/domain/repository"
)

type SettingsI interface {
	Get() (*entity.Settings, error)
	Update(settings *entity.Settings) error
}

type Settings struct {
	settingsRepo repository.Settings
}

func (s Settings) Get() (*entity.Settings, error) {
	return s.settingsRepo.Get()
}

func (s Settings) Update(settings *entity.Settings) error {
	return s.settingsRepo.Update(settings)
}

func NewSettings(settingsRepo repository.Settings) *Settings {
	return &Settings{settingsRepo}
}

var _ SettingsI = &Settings{}
