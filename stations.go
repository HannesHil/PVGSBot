package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

type Stop struct {
	stopID   string
	stopName string
	stopLon  string
	stopLat  string
}

var stopMap map[string]Stop

func loadStations() {

	csvfile, err := os.Open("PVGSStops.txt")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))
	stopMap = make(map[string]Stop)
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		stopMap[strings.ToLower(record[1])] = Stop{record[0], record[1], record[2], record[3]}
	}

	log.Println(len(stopMap))
}

func findStopByName(stopNames []string) []Stop {
	foundNames := make([]Stop, 0)
	for k, v := range stopMap {
		contained := true
		for stopNamePart := range stopNames {
			if strings.Contains(k, strings.ToLower(stopNames[stopNamePart])) {
			} else {
				contained = false
			}
		}
		if contained {
			foundNames = append(foundNames, v)
		}
	}
	return foundNames
}
