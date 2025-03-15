package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/goccy/go-json"
	"github.com/misnaged/scriptorium/logger"
	"io"
	"net/http"
	"postgres-reforger/internal/models"
	"postgres-reforger/internal/service"
	"strings"
)

type IServer interface {
	Route()
	Serve()
	Stop()
}
type HTTPServer struct {
	srvc service.IService
	*http.Server
}

func NewServer(srvc service.IService) IServer {
	srv := &HTTPServer{
		Server: &http.Server{
			Addr: ":8080",
		},
	}
	srv.srvc = srvc
	return srv
}

func (s *HTTPServer) PrepareChimChar(w http.ResponseWriter, r *http.Request) {
	if err := s.srvc.PrepareChimeraCharacters(); err != nil {
		logger.Log().Errorf("%v", err)
		return
	}
}

func (s *HTTPServer) GetUUID(w http.ResponseWriter, r *http.Request) {
	stri := r.URL.Query().Get("username")
	toWrite, err := s.srvc.GetUUIDFromName(stri)
	if err != nil {
		logger.Log().Errorf("%v", err)
		return
	}
	//str := fmt.Sprintf("%v\n", toWrite)
	//fmt.Printf("towrite:%s str:%s ", toWrite, str)
	fmt.Println(toWrite.String())
	w.Write([]byte(toWrite.String()))
}

func (s *HTTPServer) UpdateWeather(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log().Errorf("%v", err)
		return
	}
	var timeWeather models.TimeWeather
	if err = json.Unmarshal(b, &timeWeather); err != nil {
		logger.Log().Errorf("character unmarshalling failed %v", err)
		return
	}

	if err = s.srvc.CreateOrUpdateWeatherTime(&timeWeather); err != nil {
		logger.Log().Errorf("weather updating failed %v", err)
		return
	}
}
func (s *HTTPServer) GetWeather(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log().Errorf("%v", err)
	}
	conditionData, err := models.NewCondition(b)
	if err != nil {
		logger.Log().Errorf("%v", err)
	}

	toWrite, err := s.srvc.GetWeatherTime(conditionData.Condition.ComparisonValues[0])
	if err != nil {
		logger.Log().Errorf("%v", err)
		return
	}

	_, err = w.Write(toWrite)
	if err != nil {
		logger.Log().Errorf("%v", err)
		return
	}
}
func (s *HTTPServer) GetRootEntity(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log().Errorf("%v", err)
	}
	conditionData, err := models.NewCondition(b)
	if err != nil {
		logger.Log().Errorf("%v", err)
	}

	fmt.Println(conditionData.Condition.ComparisonValues[0])
	toWrite, err := s.srvc.GetRootEntity(conditionData.Condition.ComparisonValues[0])
	if err != nil {
		logger.Log().Errorf("%v", err)
		return
	}

	_, err = w.Write(toWrite)
	if err != nil {
		logger.Log().Errorf("%v", err)
		return
	}
	fmt.Println("get", string(toWrite))
}
func (s *HTTPServer) UpdateRootEntity(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log().Errorf("%v", err)
		return
	}
	var root models.RootEntity
	if err = json.Unmarshal(b, &root); err != nil {
		logger.Log().Errorf("character unmarshalling failed %v", err)
		return
	}

	if err = s.srvc.CreateOrUpdateRootEntity(&root); err != nil {
		logger.Log().Errorf("root entity collection updating failed %v", err)
		return
	}
}

func (s *HTTPServer) GetCharacter(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log().Errorf("%v", err)
	}

	conditionData, err := models.NewCondition(b)
	if err != nil {
		logger.Log().Errorf("%v", err)
	}

	fmt.Println(conditionData.Condition.ComparisonValues[0])
	toWrite, err := s.srvc.GetCharacter(conditionData.Condition.ComparisonValues[0])
	if err != nil {
		logger.Log().Errorf("%v", err)
		return
	}

	_, err = w.Write(toWrite)
	if err != nil {
		logger.Log().Errorf("%v", err)
		return
	}
	fmt.Println("get", string(toWrite))
}
func (s *HTTPServer) UpdateCharacter(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log().Errorf("%v", err)
		return
	}
	var character models.Character
	if err = json.Unmarshal(b, &character); err != nil {
		logger.Log().Errorf("character unmarshalling failed %v", err)
		return
	}

	var jsonDataMap map[string]interface{}
	if err := json.Unmarshal(b, &jsonDataMap); err != nil {
		logger.Log().Errorf("character updating failed %v", err)
		return
	}

	result := make(map[string]string)
	models.ParseJSON(jsonDataMap, result)

	for id, name := range result {
		if err = s.srvc.CreateEntHelp(id, name); err != nil {
			logger.Log().Errorf("character updating failed %v", err)
			return
		}
	}

	if err = s.srvc.CreateOrUpdateCharacter(&character); err != nil {
		logger.Log().Errorf("character updating failed %v", err)
		return
	}
}

func (s *HTTPServer) DeleteCharacter(w http.ResponseWriter, r *http.Request) {
	stri := r.URL.Query().Get("char_uuid")

	if err := s.srvc.RemoveCharacter(stri); err != nil {
		logger.Log().Errorf("removing character failed %v", err)
		return
	}
}
func (s *HTTPServer) DeleteWeather(w http.ResponseWriter, r *http.Request) {
	stri := r.URL.Query().Get("uuid")

	if err := s.srvc.RemoveWeatherTime(stri); err != nil {
		logger.Log().Errorf("removing Weather failed %v", err)
		return
	}
}
func (s *HTTPServer) DeleteRootEntity(w http.ResponseWriter, r *http.Request) {
	stri := r.URL.Query().Get("uuid")

	if err := s.srvc.RemoveRootEntityCollection(stri); err != nil {
		logger.Log().Errorf("removing RootEntity failed %v", err)
		return
	}
}
func (s *HTTPServer) DeleteItem(w http.ResponseWriter, r *http.Request) {
	stri := r.URL.Query().Get("uuid")

	if err := s.srvc.RemoveItem(stri); err != nil {
		logger.Log().Errorf("removing Item's failed %v", err)
		return
	}
}
func (s *HTTPServer) UpdateItem(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log().Errorf("%v", err)
		return
	}

	var item models.Item
	if err = json.Unmarshal(b, &item); err != nil {
		logger.Log().Errorf("item unmarshalling failed %v", err)
		return
	}
	var jsonDataMap map[string]interface{}
	if err := json.Unmarshal(b, &jsonDataMap); err != nil {
		logger.Log().Errorf("character updating failed %v", err)
		return
	}

	result := make(map[string]string)
	models.ParseJSON(jsonDataMap, result)

	for id, name := range result {
		if err = s.srvc.CreateEntHelp(id, name); err != nil {
			logger.Log().Errorf("character updating failed %v", err)
			return
		}
	}
	if err = s.srvc.CreateOrUpdateItem(&item); err != nil {
		logger.Log().Errorf("character updating failed %v", err)
		return
	}
}

func ConvertString(s string) string {
	strOut := fmt.Sprintf("%s%s", s, "_DIVIDER")
	return strOut
}
func (s *HTTPServer) GetItem(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Log().Errorf("%v", err)
	}
	conditionData, err := models.NewCondition(b)
	if err != nil {
		logger.Log().Errorf("%v", err)
		return
	}
	var kek []*models.Item

	fmt.Println("compar val is", len(conditionData.Condition.ComparisonValues))
	if len(conditionData.Condition.ComparisonValues) == 1 {
		fmt.Println("ComparisonValues", conditionData.Condition.ComparisonValues[0])
		itemJson, err, ok := s.srvc.GetItem(conditionData.Condition.ComparisonValues[0])
		if err != nil {
			logger.Log().Errorf("%v", err)
			return
		}
		if ok {

			toWrite, err := json.Marshal(itemJson)
			if err != nil {
				logger.Log().Errorf("%v", err)
				return
			}
			_, err = w.Write(toWrite)
			if err != nil {
				logger.Log().Errorf("%v", err)
				return
			}
			fmt.Println("get", string(toWrite))
		}
	} else {
		for _, v := range conditionData.Condition.ComparisonValues {
			fmt.Println("condition:", v)
			itemJson, err, ok := s.srvc.GetItem(v)
			if err != nil {
				logger.Log().Errorf("%v", err)
				return
			}
			if ok {
				kek = append(kek, itemJson.Item)
			}

		}
		var kkk []string
		for _, v := range kek {
			toWrite, err := json.Marshal(v)
			if err != nil {
				logger.Log().Errorf("%v", err)
				return
			}

			kkk = append(kkk, string(toWrite))
		}
		var ss []string
		for _, v := range kkk {
			ss = append(ss, ConvertString(v))
		}

		sss := strings.Join(ss, "")

		_, err = w.Write([]byte(sss))
		if err != nil {
			logger.Log().Errorf("%v", err)
			return
		}
		fmt.Println("get", string([]byte(sss)))
	}
	//}

}

func (s *HTTPServer) Serve() {
	if err := s.ListenAndServe(); err != nil {
		logger.Log().Errorf("listen HTTP server: %v", err)
		return
	}
}

func (s *HTTPServer) Stop() {
	if err := s.Server.Close(); err != nil {
		logger.Log().Errorf("%v", err)
	}
}

func (s *HTTPServer) Route() {
	h := chi.NewRouter()
	h.Use(middleware.Logger)
	h.Use(middleware.Recoverer)

	h.Use(CORS())

	h.Get("/get_uuid", s.GetUUID)
	h.Get("/prepareChimChar", s.PrepareChimChar)

	h.Post("/Character", s.GetCharacter)
	h.Post("/Character/update", s.UpdateCharacter)
	h.Delete("/Character/delete", s.DeleteCharacter)

	h.Delete("/Item/delete", s.DeleteItem)
	h.Post("/Item/update", s.UpdateItem)
	h.Post("/Item", s.GetItem)

	h.Post("/RootEntityCollection", s.GetRootEntity)
	h.Post("/RootEntityCollection/update", s.UpdateRootEntity)
	h.Delete("/RootEntityCollection/delete", s.DeleteRootEntity)

	h.Post("/TimeAndWeather/update", s.UpdateWeather)
	h.Post("/TimeAndWeather", s.GetWeather)
	h.Delete("/TimeAndWeather/delete", s.DeleteWeather)

	s.Handler = h
}

// CORS is middleware that controls CORS headers
func CORS() func(next http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            false,
	})
}
