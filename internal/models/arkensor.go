package models

import (
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofrs/uuid"
)

type CharJSON struct {
	CharUUID  uuid.UUID  `json:"char_uuid"`
	Character *Character `json:"character"`
}

func NewCharJson(char *Character) *CharJSON {
	return &CharJSON{CharUUID: char.MsId}
}

type Character struct {
	MsId             uuid.UUID         `json:"m_sId"`
	MiLastSaved      int               `json:"m_iLastSaved"`
	MrPrefab         string            `json:"m_rPrefab"`
	MpTransformation *MpTransformation `json:"m_pTransformation"`
	MaComponents     []*MaComponents   `json:"m_aComponents,omitempty"`
}
type MpTransformation struct {
	MvOrigin []float64 `json:"m_vOrigin,omitempty"`
	MvAngles []float64 `json:"m_vAngles,omitempty"`
}
type MaComponents struct {
	CompType string `json:"_type"`
	Data     *Data  `json:"m_pData"`
}
type Data struct {
	LayoutVers    int          `json:"m_iDataLayoutVersion"`
	Priority      int          `json:"m_iPriority,omitempty"`
	PurposeFlag   int          `json:"m_ePurposeFlags,omitempty"`
	Slots         []*Slots     `json:"m_aSlots,omitempty"`
	QuickSlots    []*QuickSlot `json:"m_aQuickSlotEntities,omitempty"`
	AmmoCount     int          `json:"m_iAmmoCount,omitempty"`
	MuzzleType    any          `json:"m_eMuzzleType,omitempty"`
	ChamberStatus []bool       `json:"m_aChamberStatus,omitempty"`
	Hitzones      []*Hitzones  `json:"m_aHitzones,omitempty"`
}
type Hitzones struct {
	Name   string  `json:"m_sName"`
	Health float64 `json:"m_fHealth"`
}
type QuickSlot struct {
	MsIndex    int       `json:"m_iIndex"`
	MsEntityId uuid.UUID `json:"m_sEntityId"`
}
type Slots struct {
	SlotIndex  int     `json:"m_iSlotIndex"`
	SlotType   string  `json:"_type"`
	SlotEntity *Entity `json:"m_pEntity"`
}
type Entity struct {
	MsId         uuid.UUID       `json:"m_sId"`
	MiLastSaved  int             `json:"m_iLastSaved"`
	MrPrefab     string          `json:"m_rPrefab"`
	MaComponents []*MaComponents `json:"m_aComponents,omitempty"`
	EntityName   string          `json:"m_sEntityName,omitempty"`
}

type CharacterSql struct {
	tableName struct{}   `sql:"character_ark" pg:",discard_unknown_columns"`
	CharUUID  uuid.UUID  `sql:"char_uuid,pk"`
	Character *Character `sql:"character"`
}

func NewChar(char *Character) *Character {
	c := &Character{
		MsId:             char.MsId,
		MiLastSaved:      char.MiLastSaved,
		MrPrefab:         char.MrPrefab,
		MpTransformation: char.MpTransformation,
	}
	if char.MaComponents != nil {
		c.MaComponents = char.MaComponents
	}
	return c
}

func NewCharSql(char *Character) (*CharacterSql, error) {

	if char == nil {
		return nil, errors.New("character struct is nil")
	}

	charSQL := &CharacterSql{
		CharUUID:  char.MsId,
		Character: char,
	}
	return charSQL, nil
}

//**************************************//

// ConditionData is
type ConditionData struct {
	Condition *Condition `json:"condition"`
}

type Condition struct {
	FieldPath        string   `json:"fieldPath"`
	ComparisonValues []string `json:"comparisonValues"`
}

func NewCondition(b []byte) (*ConditionData, error) {
	var conditionData ConditionData
	if err := json.Unmarshal(b, &conditionData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal condition data: %w", err)
	}
	if conditionData.Condition == nil {
		return nil, errors.New("condition is nil")
	}
	if conditionData.Condition.FieldPath != "m_sId" {
		return nil, errors.New("fieldPath is not equal to m_sId")
	}
	return &conditionData, nil
}

type Node struct {
	Next       *Node
	Components []*MaComponents
	Slots      []*Slots
}

func ParseJSON(data interface{}, result map[string]string) {
	switch v := data.(type) {
	case map[string]interface{}:
		if id, ok := v["m_sId"].(string); ok {
			if name, exists := v["m_sEntityName"].(string); exists {
				result[id] = name
			}
		}
		for _, value := range v {
			ParseJSON(value, result)
		}
	case []interface{}:
		for _, item := range v {
			ParseJSON(item, result)
		}
	}
}
