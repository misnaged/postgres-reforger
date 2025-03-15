package models

import (
	"github.com/gofrs/uuid"
)

type EntityHelperSql struct {
	tableName  struct{}  `sql:"entity_helper" pg:",discard_unknown_columns"`
	Id         uuid.UUID `sql:"m_sid,pk"`
	EntityName string    `sql:"entity_name"`
}

func NewEntityHelper(id, name string) *EntityHelperSql {
	uuID, _ := uuid.FromString(id)
	entHelp := &EntityHelperSql{
		tableName:  struct{}{},
		Id:         uuID,
		EntityName: name,
	}

	return entHelp
}
