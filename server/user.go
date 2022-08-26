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
	//	start := time.Now()
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
	err = s.Cache.Set("todo", details, 1*time.Millisecond)
	if err == nil {
		fmt.Println("cache done successfully")
		return
	}

}

func (s *Server) GetData(writer http.ResponseWriter, request *http.Request) {

	//var userId models.UserId
	start := time.Now()

	var result models.ToDo
	data, err := s.Cache.Get("todo")
	if err != nil {
		logrus.Error("GetData: Error  in fetching data from  cache")
	}
	if data != nil {

		err := json.Unmarshal(data, &result)
		if err != nil {
			logrus.Error("error in unmarshal")
		}
		b, _ := json.Marshal(map[string]interface{}{
			"value":             result,
			"defaultExpiration": time.Since(start).Microseconds(),
		})

		writer.Write([]byte(b))
		return
		//time:=time.After(1*time.Minute)

	}

	fmt.Println(data)

	//result, err = helper.GetData(userId)
	//if err != nil {
	//	logrus.Error("GetData: Error in  %v", err)
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	//
	//}
	//err = utilities.Caches.Set("todo", result, 1*time.Minute)
	//utilities.Caches.Set("todo",{}
	//
	//})
	//if err != nil {
	//	logrus.Error("Error in adding data to cache")
	//}
	//
	//data, err = json.Marshal(map[string]interface{}{
	//	"data":    result,
	//	"elapsed": time.Since(start).Microseconds(),
	//})
	//
	//if err != nil {
	//	logrus.Error(err)
	//}
	//_, _ = writer.Write(data)
}

//func (s *Server) Delete() {
//	err := s.Cache.Set("todo", nil, 1*time.Minute)
//	if err == nil {
//		fmt.Println("cache erased successfully")
//		return
//	}
//}
