/*
{
  "m_sId": "00bb0001-1a79-8da7-3f92-4caf4b01cc1c",
  "m_iLastSaved": 1741372569,
  "m_sWeatherState": "Clear",
  "m_bWeatherLooping": false,
  "m_iYear": 1989,
  "m_iMonth": 12,
  "m_iDay": 29,
  "m_iHour": 7,
  "m_iMinute": 31,
  "m_iSecond": 48
}
*/

package models

import "github.com/google/uuid"

type TimeWeather struct {
	MsId         uuid.UUID `json:"m_sId"`
	MiLastSaved  int       `json:"m_iLastSaved"`
	WeatherState string    `json:"m_rPrefab"`
	Year         int       `json:"m_iYear"`
	Month        int       `json:"m_iMonth"`
	Day          int       `json:"m_iDay"`
	Hour         int       `json:"m_iHour"`
	Minute       int       `json:"m_iMinute"`
	Second       int       `json:"m_iSecond"`
}
type TimeWeatherSql struct {
	tableName   struct{}     `sql:"timeweather_ark" pg:",discard_unknown_columns"`
	UUID        uuid.UUID    `sql:"uuid,pk"`
	TimeWeather *TimeWeather `sql:"timeweather"`
}
type TimeWeatherJSON struct {
	UUID        uuid.UUID    `json:"uuid"`
	TimeWeather *TimeWeather `json:"timeweather"`
}

func NewTimeWeatherSql(t *TimeWeather) (*TimeWeatherSql, error) {

	tSql := &TimeWeatherSql{
		UUID:        t.MsId,
		TimeWeather: t,
	}
	return tSql, nil
}
