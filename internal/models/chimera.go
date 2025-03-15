package models

import (
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gofrs/uuid"
	"os"
)

type ServerOnlySQL struct {
	tableName struct{}  `sql:"server_only" pg:",discard_unknown_columns"`
	Uuid      uuid.UUID `sql:"uuid,pk"`
	Identity  string    `sql:"identity" pg:",type:text"`
	Username  string    `sql:"username" pg:",type:text"`
	FirstName string    `sql:"first_name" pg:",type:text"`
	LastName  string    `sql:"last_name" pg:",type:text"`
	SteamId   string    `sql:"steam_id" pg:",type:text"`
}

type ServerOnlyJSON struct {
	Uuid      uuid.UUID `json:"uuid"`
	Identity  string    `json:"identity"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	SteamId   string    `json:"steam_id"`
}

func PrepareChimChar() ([]*ServerOnlySQL, error) {
	var sOnlyJSON []*ServerOnlyJSON
	var sOnlySQL []*ServerOnlySQL

	b, err := os.ReadFile("..\\serv_only.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read json file chimera_character.json: %w", err)

	}

	err = json.Unmarshal(b, &sOnlyJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to Unmarshal json with chimera character: %w", err)
	}

	for i := range sOnlyJSON {
		uuidNew, err := uuid.NewV4()
		if err != nil {
			return nil, fmt.Errorf("failed to generate uuid: %w", err)
		}
		onlySql := &ServerOnlySQL{
			tableName: struct{}{},
			Uuid:      uuidNew,
			Identity:  sOnlyJSON[i].Identity,
			Username:  sOnlyJSON[i].Username,
			FirstName: sOnlyJSON[i].FirstName,
			LastName:  sOnlyJSON[i].LastName,
			SteamId:   sOnlyJSON[i].SteamId,
		}
		sOnlySQL = append(sOnlySQL, onlySql)
	}
	sOnlySQLBytes, err := json.Marshal(sOnlySQL)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ServerOnly item %w", err)
	}

	if err = json.Unmarshal(sOnlySQLBytes, &sOnlySQL); err != nil {
		return nil, fmt.Errorf("failed to UNmarshal ChimeraCharacter item %w", err)
	}

	return sOnlySQL, nil
}

//func NewChimCharSql(chimchar *ChimeraCharacterJSON) (*ChimeraCharacterSQL, error) {
//
//	chimcharSql := &ChimeraCharacterSQL{
//		tableName: struct{}{},
//		Uuid:      chimchar.Uuid,
//		Identity:  chimchar.Identity,
//		Username:  chimchar.Username,
//		FirstName: chimchar.FirstName,
//		LastName:  chimchar.LastName,
//		SteamId:   chimchar.SteamId,
//		Pos:       chimchar.Pos,
//	}
//
//	b, err := json.Marshal(chimcharSql)
//	if err != nil {
//		return nil, fmt.Errorf("failed to marshal ChimeraCharacter item %w", err)
//	}
//
//	if err = json.Unmarshal(b, &chimcharSql); err != nil {
//		return nil, fmt.Errorf("failed to UNmarshal ChimeraCharacter item %w", err)
//	}
//
//	return chimcharSql, nil
//}
