package service

import (
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"postgres-reforger/config"
	"postgres-reforger/internal/models"
	"postgres-reforger/internal/repository"
)

type IService interface {
	PrepareChimeraCharacters() error
	GetUUIDFromName(name string) (*uuid.UUID, error)

	CreateOrUpdateCharacter(model *models.Character) error
	GetCharacter(param string) ([]byte, error)
	RemoveCharacter(param string) error

	CreateOrUpdateItem(model *models.Item) error
	GetItem(param string) (*models.ItemJSON, error, bool)
	RemoveItem(param string) error

	CreateOrUpdateRootEntity(model *models.RootEntity) (err error)
	GetRootEntity(param string) ([]byte, error)
	RemoveRootEntityCollection(param string) error

	CreateOrUpdateWeatherTime(model *models.TimeWeather) error
	GetWeatherTime(param string) ([]byte, error)
	RemoveWeatherTime(param string) error

	CreateEntHelp(id, name string) error
}

type ReforgerDbService struct {
	cfg  *config.Scheme
	repo repository.IRepository
}

func NewService(cfg *config.Scheme, repo repository.IRepository) IService {
	return &ReforgerDbService{cfg: cfg, repo: repo}
}

func (rfg *ReforgerDbService) CreateOrUpdateWeatherTime(model *models.TimeWeather) error {
	timeSQL, err := models.NewTimeWeatherSql(model)
	if err != nil {
		return fmt.Errorf("unable to create new weather model: %w", err)
	}
	if err = rfg.repo.CreateOrUpdateWeatherTime(timeSQL); err != nil {
		return fmt.Errorf("updating weather and time storage failed: %w", err)
	}
	return nil
}

func (rfg *ReforgerDbService) GetWeatherTime(param string) ([]byte, error) {
	timeSQL, err := rfg.repo.SelectWeatherTime(param)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	timeJSON := &models.TimeWeatherJSON{TimeWeather: timeSQL[0].TimeWeather}
	timeJSON.UUID = timeJSON.TimeWeather.MsId

	b, err := json.Marshal(timeJSON.TimeWeather)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal %w", err)
	}
	return b, nil
}

func (rfg *ReforgerDbService) RemoveCharacter(param string) error {
	_, err := rfg.repo.RemoveCharacter(param)
	if err != nil {
		return fmt.Errorf("failed to Remove Character from repo %w", err)
	}
	return nil
}

func (rfg *ReforgerDbService) RemoveRootEntityCollection(param string) error {
	_, err := rfg.repo.RemoveRootEntityCollection(param)
	if err != nil {
		return fmt.Errorf("failed to Remove RootEntityCollection from repo %w", err)
	}
	return nil
}

func (rfg *ReforgerDbService) RemoveWeatherTime(param string) error {
	_, err := rfg.repo.RemoveWeatherTime(param)
	if err != nil {
		return fmt.Errorf("failed to Remove Weather from repo %w", err)
	}
	return nil
}

func (rfg *ReforgerDbService) CreateEntHelp(id, name string) error {
	entSql := models.NewEntityHelper(id, name)
	if err := rfg.repo.CreateEntHelp(entSql); err != nil {
		return fmt.Errorf("creating new entity_helper failed: %w", err)
	}

	return nil
}
func (rfg *ReforgerDbService) CreateOrUpdateItem(model *models.Item) error {
	itemSql, err := models.NewItemSql(model)
	if err != nil {
		return fmt.Errorf("unable to crete new item model: %w", err)
	}
	if err = rfg.repo.CreateOrUpdateItem(itemSql); err != nil {
		return fmt.Errorf("creating new item failed: %w", err)
	}

	return nil
}
func (rfg *ReforgerDbService) GetItem(param string) (*models.ItemJSON, error, bool) {
	itemSQL, err := rfg.repo.SelectItem(param)
	if err != nil {
		return nil, fmt.Errorf("failed to SelectItem from repo %w", err), false
	}
	if itemSQL == nil {
		return nil, nil, false
	}
	return models.NewItemJSON(itemSQL[0]), nil, true
}
func (rfg *ReforgerDbService) RemoveItem(param string) error {
	_, err := rfg.repo.RemoveItem(param)
	if err != nil {
		return fmt.Errorf("failed to Remove Item from repo %w", err)
	}
	return nil
}
func (rfg *ReforgerDbService) CreateOrUpdateRootEntity(model *models.RootEntity) (err error) {
	rootSQL, err := models.NewRootSql(model)
	if err != nil {
		return fmt.Errorf("unable to update character model: %w", err)
	}
	if err = rfg.repo.CreateOrUpdateEntityCollection(rootSQL); err != nil {
		return fmt.Errorf("updating root entity storage failed: %w", err)
	}
	return nil
}
func (rfg *ReforgerDbService) GetRootEntity(param string) ([]byte, error) {
	rootSQL, err := rfg.repo.SelectRootEntityCollection(param)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	rootJSON := &models.RootEntityJSON{RootEntity: rootSQL[0].RootEntity}
	rootJSON.UUID = rootJSON.RootEntity.MsId

	b, err := json.Marshal(rootJSON.RootEntity)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal %w", err)
	}
	return b, nil
}

func (rfg *ReforgerDbService) CreateOrUpdateCharacter(model *models.Character) error {
	charSQL, err := models.NewCharSql(model)
	if err != nil {
		return fmt.Errorf("unable to update character model: %w", err)
	}
	if err = rfg.repo.CreateOrUpdateCharacter(charSQL); err != nil {
		return fmt.Errorf("updating character storage failed: %w", err)
	}
	return nil
}
func (rfg *ReforgerDbService) GetCharacter(param string) ([]byte, error) {
	charSQL, err := rfg.repo.SelectCharacterSql(param)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	charJSON := &models.CharJSON{Character: charSQL[0].Character}
	charJSON.CharUUID = charJSON.Character.MsId

	b, err := json.Marshal(charJSON.Character)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal %w", err)
	}
	return b, nil
}

func (rfg *ReforgerDbService) GetUUIDFromName(name string) (*uuid.UUID, error) {
	aa, err := rfg.repo.GetChimeraCharByName(name)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	return &aa[0].Uuid, nil
}

func removeString(slice []string, strToRemove string) []string {
	var result []string
	for _, v := range slice {
		if v != strToRemove {
			result = append(result, v)
		}
	}
	return result
}

func (rfg *ReforgerDbService) PrepareChimeraCharacters() error {
	chimchars, err := models.PrepareChimChar()
	if err != nil {
		return fmt.Errorf("unable to crete new chimchar model: %w", err)
	}

	for i := range chimchars {
		if err = rfg.repo.CreateChimera(chimchars[i]); err != nil {
			return fmt.Errorf("creating new chimchar failed: %w", err)
		}
	}

	return nil
}
