package db

import (
	"fmt"
	"log"
)

func validateFilterLingsRequest(flr FilterLingsRequest) error {
	hasGroup := flr.Group != 0
	hasLings := len(flr.Lings) != 0

	if hasGroup && hasLings {
		return nil
	}

	log.Printf("Malformed request: %+v", flr)
	return fmt.Errorf("Malformed filter request!")
}

func validateFilterLingPropertiesRequest(flpr FilterLingPropertiesRequest) error {
	hasGroup := flpr.Group != 0
	hasLingProperties := len(flpr.LingProperties) != 0

	if hasGroup && hasLingProperties {
		return nil
	}

	log.Printf("Malformed request: %+v", flpr)
	return fmt.Errorf("Malformed filter request!")
}

func validateFilterLingletsRequest(fllr FilterLingletsRequest) error {
	hasGroup := fllr.Group != 0
	hasLinglets := len(fllr.Linglets) != 0

	if hasGroup && hasLinglets {
		return nil
	}

	log.Printf("Malformed request: %+v", fllr)
	return fmt.Errorf("Malformed filter request!")
}

func validateFilterLingletPropertiesRequest(fllpr FilterLingletPropertiesRequest) error {
	hasGroup := fllpr.Group != 0
	hasLingletProperties := len(fllpr.LingletProperties) != 0

	if hasGroup && hasLingletProperties {
		return nil
	}

	log.Printf("Malformed request: %+v", fllpr)
	return fmt.Errorf("Malformed filter request!")
}

func validateCompareLingsRequest(clr CompareLingsRequest) error {
	hasGroup := clr.Group != 0
	hasLings := len(clr.Lings) != 0
	leqSix := len(clr.Lings) <= 6

	if hasGroup && hasLings && leqSix {
		return nil
	}

	log.Printf("Malformed request: %+v", clr)
	return fmt.Errorf("Malformed compare request!")
}

func validateCompareLingletsRequest(cllr CompareLingletsRequest) error {
	hasGroup := cllr.Group != 0
	hasLinglets := len(cllr.Linglets) != 0
	leqSix := len(cllr.Linglets) <= 6

	if hasGroup && hasLinglets && leqSix {
		return nil
	}

	log.Printf("Malformed request: %+v", cllr)
	return fmt.Errorf("Malformed compare request!")
}

func validateCrossLingPropertiesRequest(clpr CrossLingPropertiesRequest) error {
	hasGroup := clpr.Group != 0
	hasLingProperties := len(clpr.LingProperties) != 0
	leqSix := len(clpr.LingProperties) <= 6

	if hasGroup && hasLingProperties && leqSix {
		return nil
	}

	log.Printf("Malformed request: %+v", clpr)
	return fmt.Errorf("Malformed cross request!")
}

func validateCrossLingletPropertiesRequest(cllpr CrossLingletPropertiesRequest) error {
	hasGroup := cllpr.Group != 0
	hasLingletProperties := len(cllpr.LingletProperties) != 0
	leqSix := len(cllpr.LingletProperties) <= 6

	if hasGroup && hasLingletProperties && leqSix {
		return nil
	}

	log.Printf("Malformed request: %+v", cllpr)
	return fmt.Errorf("Malformed cross request!")
}
