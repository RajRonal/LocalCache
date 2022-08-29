package server

import (
	"InMemoryCache/database/helper"
	"InMemoryCache/models"
	"fmt"

	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func (s *Server) InsertData(writer http.ResponseWriter, request *http.Request) {
	var details models.ToDo
	err := json.NewDecoder(request.Body).Decode(&details)
	if err != nil {
		logrus.Error("InsertData: Error in decoding json %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = helper.InsertData(details)
	if err != nil {
		logrus.Error("InsertData: Error in Inserting data %v", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	//fmt.Println(time.Now().Add(2 * time.Minute))
	_, err = s.Cache.Set("todo", details, time.Now().Add(30*time.Second))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

}

func (s *Server) GetData(writer http.ResponseWriter, request *http.Request) {

	var outputData models.ToDo
	dbData, err := s.Cache.Get("todo")
	if err != nil {
		logrus.Error("GetData: Error  in fetching dbData from  cache")
		writer.Write([]byte("No data In Cache"))
	}
	if dbData != nil {

		errs := json.Unmarshal(dbData, &outputData)
		if errs != nil {
			logrus.Error("error in unmarshal")
			writer.WriteHeader(http.StatusBadRequest)
		}
		//

		b, _ := json.Marshal(map[string]interface{}{
			"value":             outputData,
			"defaultExpiration": outputData.ExpiryTime,
		})
		writer.Write([]byte(b))
		return

	}

	fmt.Println(dbData)

}
