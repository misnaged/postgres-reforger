package repository

import (
	"fmt"
	"postgres-reforger/internal/models"
)

func (r *Repository) CreateChimera(model *models.ServerOnlySQL) (err error) {
	_, err = r.db.Model(model).Insert()
	if err != nil {
		return fmt.Errorf("insert chimera characters into db: %w", err)
	}
	return nil

}

func (r *Repository) GetChimeraCharByName(name string) (model []*models.ServerOnlySQL, err error) {
	query := r.db.Model(&model).Column("uuid").Where("username =?", name)
	if err := query.Select(); err != nil {
		return nil, fmt.Errorf("select uuid from database failed: %w", err)
	}
	return model, nil
}
