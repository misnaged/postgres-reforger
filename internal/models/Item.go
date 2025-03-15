/*
{
  "m_sId": "00bb0001-5cbd-89be-110b-f707755dc2fb",
  "m_iLastSaved": 1741372569,
  "m_rPrefab": "E65D5C9848AC18E9",
  "m_pTransformation": {
    "m_vOrigin": [
      1232.24865,
      152.36842,
      1892.18457
    ],
    "m_vAngles": [
      -128.70698,
      0,
      0
    ]
  },
  "m_fRemainingLifetime": 58.734
}
*/

package models

import (
	"errors"
	"github.com/gofrs/uuid"
)

type Item struct {
	MsId              uuid.UUID         `json:"m_sId"`
	MiLastSaved       int               `json:"m_iLastSaved"`
	MrPrefab          string            `json:"m_rPrefab"`
	RemainingLifetime float64           `json:"m_fRemainingLifetime"`
	MpTransformation  *MpTransformation `json:"m_pTransformation,omitempty"`
	MaComponents      []*MaComponents   `json:"m_aComponents,omitempty"`
	EntityName        string            `json:"m_sEntityName"`
}
type ItemSql struct {
	tableName struct{}  `sql:"item_ark" pg:",discard_unknown_columns"`
	UUID      uuid.UUID `sql:"uuid,pk"`
	Item      *Item     `sql:"item"`
}
type ItemJSON struct {
	UUID uuid.UUID `json:"uuid"`
	Item *Item     `json:"item"`
}

func NewItemJSON(sql *ItemSql) *ItemJSON {
	return &ItemJSON{
		UUID: sql.UUID,
		Item: sql.Item,
	}
}
func NewItemSql(item *Item) (*ItemSql, error) {

	if item == nil {
		return nil, errors.New("item struct is nil")
	}

	itemSQL := &ItemSql{
		UUID: item.MsId,
		Item: item,
	}
	return itemSQL, nil
}
