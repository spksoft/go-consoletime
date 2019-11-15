package gconsoletime

import (
	"log"
	"time"
)

var ct *consoleTime

type consoleTime struct {
	loggedTime map[string]time.Time
}

// Log is a function push and pop time of called
func Log(name string) {
	if start, ok := ct.loggedTime[name]; ok {
		elapsed := time.Since(start)
		log.Printf("%s took %s", name, elapsed)
		delete(ct.loggedTime, name)
	} else {
		start := time.Now()
		ct.loggedTime[name] = start
	}
}

func init() {
	ct = new()
}

func new() *consoleTime {
	return &consoleTime{
		loggedTime: make(map[string]time.Time),
	}
}
