package db

import (
	"fmt"
	"log"
)

func validateFilterRequest(frq FilterRequest) error {
	hasLing := len(frq.Lings) != 0
	hasLingProperties := len(frq.LingProperties) != 0
	hasLinglet := len(frq.Linglets) != 0
	hasLingletProperties := len(frq.LingletProperties) != 0

	if hasLing || hasLingProperties || hasLinglet || hasLingletProperties {
		return nil
	}

	log.Printf("Malformed request: %+v", frq)

	return fmt.Errorf("Missing field from filter request!")
}