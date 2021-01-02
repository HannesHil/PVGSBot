package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Telegram struct {
		Bottoken string
	}
	Bot struct {
		Restbaseurl string
	}
}

var B *tb.Bot
var C Config

func main() {
	f, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("Read Config & Creating Bot")
	Bot, err := tb.NewBot(tb.Settings{
		Token:  cfg.Telegram.Bottoken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	B = Bot
	fmt.Println(cfg)
	C = cfg
	fmt.Println(C)
	if err != nil {
		log.Fatal(err)
		return
	}

	B.Handle("/start", startHandler)
	B.Handle("/abfahrten", departuresHandler)

	log.Printf("Starting Bot")

	loadStations()
	B.Start()
}

func startHandler(M *tb.Message) {
	B.Send(M.Sender, "Du kannst die Abfahrten einer PVGS Station mithilfe von /abfahrten [Stationsname] aufrufen.")
}

func departuresHandler(m *tb.Message) {
	requestedStop := strings.Trim(strings.Replace(strings.ToLower(m.Text), "/abfahrten", "", 1), " ")
	requestedStopParts := strings.Split(requestedStop, " ")
	foundStops := findStopByName(requestedStopParts)
	var message string
	if len(foundStops) > 10 {
		message = fmt.Sprintf("%v Stops mit diesem Namen gefunden. Bitte schränke deine Suche ein.", len(foundStops))
	} else if len(foundStops) == 0 {
		message = "Keine Stop mit diesem Namen gefunden. Verallgemeinere deine Suche"
	} else if len(foundStops) == 1 {
		message = "Suche nach den Abfahrten für " + foundStops[0].stopName + ":\n"
		message = message + getDepartuesForStopString(C.Bot.Restbaseurl, foundStops[0])
	} else {
		message = "Folgende Stops gefunden. Bitte sende noch eine Anfrage mit dem genauen Namen: \n"
		for localStop := range foundStops {
			message = message + "- " + foundStops[localStop].stopName + "\n"
		}
	}
	B.Send(m.Sender, message, tb.ModeMarkdown)
}
