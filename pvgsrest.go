package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Departure struct {
	tripId      string
	pos         Stop
	producttype string
	productname string
	when        string
	whenplanned string
	operator    string
	delay       string
}

func getDepartuesForStop(baseURL string, stationID Stop) []map[string]interface{} {
	response, err := http.Get("https://" + baseURL + "/departures?stopID=" + stationID.stopID)
	mapStructure := make([]map[string]interface{}, 1, 1)
	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		err2 := json.Unmarshal(data, &mapStructure)
		if err2 != nil {
			log.Print(err2)
		}
		return mapStructure
	}
	return mapStructure
}

func getDepartuesForStopString(baseURL string, stationID Stop) string {
	sliceofMapofDepartures := getDepartuesForStop(baseURL, stationID)
	msg := ""
	for _, mapOfDeparts := range sliceofMapofDepartures {
		msg = msg + "*" + mapOfDeparts["line"].(map[string]interface{})["name"].(string) + "* ➔" + mapOfDeparts["direction"].(string)
		parsedTime, err2 := time.Parse(time.RFC3339, mapOfDeparts["when"].(string))
		if err2 != nil {
			log.Print(err2)
		}
		arival := parsedTime.Sub(time.Now()).Minutes()
		msg = msg + " in " + strconv.FormatInt(int64(arival), 10) + "min 🕐" + parsedTime.Format("15:04") + "\n"
	}
	return msg
}
