package webscraperscrapeconfig

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
)

type Profile string

var scraperConfigInvalidFormatMsg = func(msg string) string { return fmt.Sprintf("scraper config invalid format: %v", msg) }

func (p *Profile) Name() string {
	var m map[string]interface{}

	err := json.Unmarshal([]byte(*p), &m)
	if err != nil {
		log.Fatalf("could not unmarshal profile: %v", err)
	}

	if _, ok := m["_id"]; !ok {
		log.Fatalf(scraperConfigInvalidFormatMsg("_id field not found"))
	}

	if _, ok := m["_id"].(string); !ok {
		log.Fatalf(scraperConfigInvalidFormatMsg("_id field not string"))
	}

	return m["_id"].(string)
}

var Profiles = []Profile{
	Profile(NYCCommutableHomes),
}

//go:embed NYC-commutable-homes-with-land.json
var NYCCommutableHomes string
