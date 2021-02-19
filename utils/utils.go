package utils

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

// Load returns poses from a specified .json file
func Load(filePath string) ([]string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Errorf("Error loading .json file: %s", err)
		return nil, err
	}

	var responses []string
	err = json.Unmarshal(content, &responses)
	if err != nil {
		log.Errorf("Error unmarshalling json: %s", err)
		return nil, err
	}

	return responses, nil
}

// SeedRNG calls rand.Seed()
func SeedRNG() {
	rand.Seed(time.Now().UnixNano())
}
