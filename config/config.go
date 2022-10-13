package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	m "github.com/jovanfrandika/livechart-notifier/domain"
)

func Init() (cfg *m.Config) {
	cfg = &m.Config{}
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}

	bts, err := ioutil.ReadFile(filepath.Join(basePath, "files", "config.json"))
	if err != nil {
		log.Fatal(err.Error())
	}

	err = json.Unmarshal(bts, cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	return cfg
}
