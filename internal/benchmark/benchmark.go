package benchmark

import (
	"log"
	"time"
)

var benchmarks map[string]time.Time
var tare time.Duration

// Start - starts the timer for a specified string
func Start(t string) {
	if benchmarks == nil {
		benchmarks = make(map[string]time.Time)
	}

	log.Printf("Timer '%v' started", t)
	benchmarks[t] = time.Now()
}

// Stop - stops the timer for a specified string
func Stop(t string) {
	u := time.Now().Add(-tare)

	if b, ok := benchmarks[t]; ok {
		log.Printf("Timer '%v' stopped with time: %v", t, u.Sub(b))
	} else {
		log.Printf("No such timer found: %v", t)
	}
}

// StartTare - starts timer for tare
func StartTare() {
	if benchmarks == nil {
		benchmarks = make(map[string]time.Time)
	}

	log.Print("Tare timer started")
	benchmarks["tare"] = time.Now()
}

// StopTare - stops the timer for tare and sets tare value
func StopTare() {
	u := time.Now()

	if b, ok := benchmarks["tare"]; ok {
		tare = u.Sub(b)
		log.Printf("Tare timer stopped with time: %v", tare)
	} else {
		log.Printf("No tare timer found")
	}
}

// ResetTare - sets tare to zero
func ResetTare() {
	tare = time.Duration(0)
}
