package main

import (
	"log"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Telegram struct {
		Bottoken string
	}
}

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
	b, err := tb.NewBot(tb.Settings{
		Token:  cfg.Telegram.Bottoken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hello World!")
	})

	log.Printf("Starting Bot")
	b.Start()
}
