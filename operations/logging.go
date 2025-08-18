package operations

import (
	"log"
	"time"
)

func logStatus(status any) {
	log.Printf("%v\n", status)
}

func logExchange(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration) {
	log.Printf("%v %v %v %v %v %v\n", start, duration, route, req, resp, timeout)
}
