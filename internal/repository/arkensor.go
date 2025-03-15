package repository

import (
	"fmt"
	"github.com/gofrs/uuid"
	"postgres-reforger/internal/models"
)

func (r *Repository) CreateOrUpdateCharacter(model *models.CharacterSql) (err error) {
	_, err = r.db.Model(model).OnConflict("(char_uuid) DO UPDATE").Insert()

	if err != nil {
		return fmt.Errorf("insert character into db: %w", err)
	}
	return nil
}

func (r *Repository) SelectCharacterSql(param string) (model []*models.CharacterSql, err error) {
	uuID, _ := uuid.FromString(param)
	query := r.db.Model(&model).Column("character").Where("char_uuid =?", uuID)

	if err = query.Select(); err != nil {
		return nil, fmt.Errorf("select character from database %w", err)
	}
	return model, nil
}
func (r *Repository) RemoveCharacter(param string) (model []*models.CharacterSql, err error) {
	query := r.db.Model(&model).Column("character").Where("char_uuid =?", param)

	if _, err = query.Delete(); err != nil {
		return nil, fmt.Errorf("delete character from database %w", err)
	}
	return nil, nil
}
func (r *Repository) CreateOrUpdateEntityCollection(model *models.RootEntitySql) (err error) {
	_, err = r.db.Model(model).OnConflict("(uuid) DO UPDATE").Insert()

	if err != nil {
		return fmt.Errorf("insert RootEntityCollection into db: %w", err)
	}
	return nil
}

func (r *Repository) SelectRootEntityCollection(param string) (model []*models.RootEntitySql, err error) {
	uuID, _ := uuid.FromString(param)
	query := r.db.Model(&model).Column("root_entity").Where("uuid =?", uuID)

	if err = query.Select(); err != nil {
		return nil, fmt.Errorf("select root_entity from database %w", err)
	}
	return model, nil
}
func (r *Repository) RemoveRootEntityCollection(param string) (model []*models.RootEntitySql, err error) {
	query := r.db.Model(&model).Column("root_entity").Where("uuid =?", param)

	if _, err = query.Delete(); err != nil {
		return nil, fmt.Errorf("delete root_entity from database %w", err)
	}
	return nil, nil
}

func (r *Repository) CreateOrUpdateWeatherTime(model *models.TimeWeatherSql) (err error) {
	_, err = r.db.Model(model).OnConflict("(uuid) DO UPDATE").Insert()
	if err != nil {
		return fmt.Errorf("insert timeweather into db: %w", err)
	}
	return nil
}

func (r *Repository) SelectWeatherTime(param string) (model []*models.TimeWeatherSql, err error) {
	uuID, _ := uuid.FromString(param)
	query := r.db.Model(&model).Column("timeweather").Where("uuid =?", uuID)

	if err = query.Select(); err != nil {
		return nil, fmt.Errorf("select character from database %w", err)
	}
	return model, nil
}
func (r *Repository) RemoveWeatherTime(param string) (model []*models.TimeWeatherSql, err error) {
	query := r.db.Model(&model).Column("timeweather").Where("uuid =?", param)

	if _, err = query.Delete(); err != nil {
		return nil, fmt.Errorf("delete root_entity from database %w", err)
	}
	return nil, nil
}

func (r *Repository) CreateOrUpdateItem(model *models.ItemSql) (err error) {
	_, err = r.db.Model(model).OnConflict("(uuid) DO UPDATE").Insert()

	if err != nil {
		return fmt.Errorf("insert Item into db: %w", err)
	}
	return nil
}

func (r *Repository) SelectItem(param string) (model []*models.ItemSql, err error) {
	uuID, _ := uuid.FromString(param)
	query := r.db.Model(&model).Column("item").Where("uuid =?", uuID)

	if err = query.Select(); err != nil {
		return nil, fmt.Errorf("select character from database %w", err)
	}
	return model, nil
}

func (r *Repository) RemoveItem(param string) (model []*models.ItemSql, err error) {
	query := r.db.Model(&model).Column("item").Where("uuid =?", param)

	if _, err = query.Delete(); err != nil {
		return nil, fmt.Errorf("delete item from database %w", err)
	}
	return nil, nil
}

func (r *Repository) CreateEntHelp(model *models.EntityHelperSql) (err error) {
	_, err = r.db.Model(model).OnConflict("(m_sid) DO UPDATE").Insert()

	if err != nil {
		return fmt.Errorf("insert entity_helper into db: %w", err)
	}
	return nil
}
