package db

import (
	"fmt"
	"log"
)

func validateFilterRequest(frq FilterRequest) error {
	hasGroup := frq.Group != 0
	hasLing := len(frq.Lings) != 0
	hasLingProperties := len(frq.LingProperties) != 0
	hasLinglet := len(frq.Linglets) != 0
	hasLingletProperties := len(frq.LingletProperties) != 0

	// needs group, and one of either (ling + properties) or (linglet + properties)
	if hasGroup && ((hasLing || hasLingProperties) != (hasLinglet || hasLingletProperties)) {
		return nil
	}

	log.Printf("Malformed request: %+v", frq)

	return fmt.Errorf("Malformed filter request!")
}
