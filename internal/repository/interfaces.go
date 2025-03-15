package repository

import (
	"postgres-reforger/internal/models"
)

type IRepository interface {
	CreateChimera(model *models.ServerOnlySQL) (err error)
	GetChimeraCharByName(name string) (model []*models.ServerOnlySQL, err error)
	//SelectCharSql(param string) (model []*models.ServerOnlySQL, err error)

	CreateOrUpdateCharacter(model *models.CharacterSql) (err error)
	SelectCharacterSql(param string) (model []*models.CharacterSql, err error)
	RemoveCharacter(param string) (model []*models.CharacterSql, err error)

	CreateOrUpdateItem(model *models.ItemSql) (err error)
	SelectItem(param string) (model []*models.ItemSql, err error)
	RemoveItem(param string) (model []*models.ItemSql, err error)

	CreateOrUpdateEntityCollection(model *models.RootEntitySql) (err error)
	SelectRootEntityCollection(param string) (model []*models.RootEntitySql, err error)
	RemoveRootEntityCollection(param string) (model []*models.RootEntitySql, err error)

	CreateOrUpdateWeatherTime(model *models.TimeWeatherSql) (err error)
	SelectWeatherTime(param string) (model []*models.TimeWeatherSql, err error)
	RemoveWeatherTime(param string) (model []*models.TimeWeatherSql, err error)

	CreateEntHelp(model *models.EntityHelperSql) (err error)
}
