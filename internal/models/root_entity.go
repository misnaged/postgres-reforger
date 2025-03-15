/*
{
  "m_sId": "00ec0001-0000-0000-0000-000000000000",
  "m_iLastSaved": 1741372569,
  "m_aRemovedBackedRootEntities": [
    "00bb0001-0f7f-0251-10b3-4de700444a7d",
    "00bb0001-7a0a-b51b-010f-b51b48a78941"
  ],
  "m_aSelfSpawnDynamicEntities": []
}
*/

package models

import (
	"errors"
	"github.com/gofrs/uuid"
)

type RootEntity struct {
	MsId                      uuid.UUID                   `json:"m_sId"`
	MiLastSaved               int                         `json:"m_iLastSaved"`
	RemovedBackedRootEntities []uuid.UUID                 `json:"m_aRemovedBackedRootEntities"`
	SelfSpawnDynamicEntities  []*SelfSpawnDynamicEntities `json:"m_aSelfSpawnDynamicEntities"`
}
type SelfSpawnDynamicEntities struct {
	SaveDataType string      `json:"m_sSaveDataType"`
	Ids          []uuid.UUID `json:"m_aIds"`
}

type RootEntitySql struct {
	tableName  struct{}    `sql:"root_entity_ark" pg:",discard_unknown_columns"`
	UUID       uuid.UUID   `sql:"uuid,pk"`
	RootEntity *RootEntity `sql:"root_entity"`
}
type RootEntityJSON struct {
	UUID       uuid.UUID   `json:"uuid"`
	RootEntity *RootEntity `json:"root_entity"`
}

func NewRootSql(root *RootEntity) (*RootEntitySql, error) {

	if root == nil {
		return nil, errors.New("character struct is nil")
	}

	rootSql := &RootEntitySql{
		UUID:       root.MsId,
		RootEntity: root,
	}
	return rootSql, nil
}
